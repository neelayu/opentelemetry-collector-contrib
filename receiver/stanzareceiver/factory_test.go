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
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"
)

func TestCreateReceiver(t *testing.T) {
	params := component.ReceiverCreateParams{
		Logger: zap.NewNop(),
	}

	t.Run("Success", func(t *testing.T) {
		factory := NewFactory(TestReceiverType{})
		receiver, err := factory.CreateLogsReceiver(context.Background(), params, factory.CreateDefaultConfig(), &mockLogsConsumer{})
		require.NoError(t, err, "receiver creation failed")
		require.NotNil(t, receiver, "receiver creation failed")
	})

	t.Run("DecodeOperatorsFailure", func(t *testing.T) {
		factory := NewFactory(TestReceiverType{})
		badCfg := factory.CreateDefaultConfig().(*TestConfig)
		badCfg.Operators = []map[string]interface{}{
			{
				"badparam": "badvalue",
			},
		}
		receiver, err := factory.CreateLogsReceiver(context.Background(), params, badCfg, &mockLogsConsumer{})
		require.Error(t, err, "receiver creation should fail if operator configs aren't valid")
		require.Nil(t, receiver, "receiver creation should fail if operator configs aren't valid")
	})
}
