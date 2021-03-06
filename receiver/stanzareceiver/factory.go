// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package stanzareceiver

import (
	"context"

	"github.com/open-telemetry/opentelemetry-log-collection/agent"
	"github.com/open-telemetry/opentelemetry-log-collection/pipeline"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver/receiverhelper"
)

// LogReceiverType is the interface used by stanza-based log receivers
type LogReceiverType interface {
	Type() configmodels.Type
	CreateDefaultConfig() configmodels.Receiver
	Decode(configmodels.Receiver) (pipeline.Config, error)
}

// NewFactory creates a factory for a Stanza-based receiver
func NewFactory(logReceiverType LogReceiverType) component.ReceiverFactory {
	return receiverhelper.NewFactory(
		logReceiverType.Type(),
		logReceiverType.CreateDefaultConfig,
		receiverhelper.WithLogs(createLogsReceiver(logReceiverType)),
	)
}

func createLogsReceiver(logReceiverType LogReceiverType) receiverhelper.CreateLogsReceiver {
	return func(
		ctx context.Context,
		params component.ReceiverCreateParams,
		cfg configmodels.Receiver,
		nextConsumer consumer.LogsConsumer,
	) (component.LogsReceiver, error) {
		pipeline, err := logReceiverType.Decode(cfg)
		if err != nil {
			return nil, err
		}

		emitter := NewLogEmitter(params.Logger.Sugar())
		logAgent, err := agent.NewBuilder(params.Logger.Sugar()).
			WithConfig(&agent.Config{Pipeline: pipeline}).
			WithDefaultOutput(emitter).
			Build()
		if err != nil {
			return nil, err
		}

		return &stanzareceiver{
			agent:    logAgent,
			emitter:  emitter,
			consumer: nextConsumer,
			logger:   params.Logger,
		}, nil
	}
}
