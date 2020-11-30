// Copyright The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/config/configmodels"
	"go.opentelemetry.io/collector/consumer/pdata"
)

// Type is the component type name.
const Type configmodels.Type = "nginxreceiver"

type metricIntf interface {
	Name() string
	New() pdata.Metric
}

// Intentionally not exposing this so that it is opaque and can change freely.
type metricImpl struct {
	name    string
	newFunc func() pdata.Metric
}

func (m *metricImpl) Name() string {
	return m.name
}

func (m *metricImpl) New() pdata.Metric {
	return m.newFunc()
}

type metricStruct struct {
	NginxConnectionsAccepted metricIntf
	NginxConnectionsActive   metricIntf
	NginxConnectionsHandled  metricIntf
	NginxConnectionsReading  metricIntf
	NginxConnectionsWaiting  metricIntf
	NginxConnectionsWriting  metricIntf
	NginxRequests            metricIntf
}

// Names returns a list of all the metric name strings.
func (m *metricStruct) Names() []string {
	return []string{
		"nginx.connections_accepted",
		"nginx.connections_active",
		"nginx.connections_handled",
		"nginx.connections_reading",
		"nginx.connections_waiting",
		"nginx.connections_writing",
		"nginx.requests",
	}
}

var metricsByName = map[string]metricIntf{
	"nginx.connections_accepted": Metrics.NginxConnectionsAccepted,
	"nginx.connections_active":   Metrics.NginxConnectionsActive,
	"nginx.connections_handled":  Metrics.NginxConnectionsHandled,
	"nginx.connections_reading":  Metrics.NginxConnectionsReading,
	"nginx.connections_waiting":  Metrics.NginxConnectionsWaiting,
	"nginx.connections_writing":  Metrics.NginxConnectionsWriting,
	"nginx.requests":             Metrics.NginxRequests,
}

func (m *metricStruct) ByName(n string) metricIntf {
	return metricsByName[n]
}

func (m *metricStruct) FactoriesByName() map[string]func() pdata.Metric {
	return map[string]func() pdata.Metric{
		Metrics.NginxConnectionsAccepted.Name(): Metrics.NginxConnectionsAccepted.New,
		Metrics.NginxConnectionsActive.Name():   Metrics.NginxConnectionsActive.New,
		Metrics.NginxConnectionsHandled.Name():  Metrics.NginxConnectionsHandled.New,
		Metrics.NginxConnectionsReading.Name():  Metrics.NginxConnectionsReading.New,
		Metrics.NginxConnectionsWaiting.Name():  Metrics.NginxConnectionsWaiting.New,
		Metrics.NginxConnectionsWriting.Name():  Metrics.NginxConnectionsWriting.New,
		Metrics.NginxRequests.Name():            Metrics.NginxRequests.New,
	}
}

// Metrics contains a set of methods for each metric that help with
// manipulating those metrics.
var Metrics = &metricStruct{
	&metricImpl{
		"nginx.connections_accepted",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("nginx.connections_accepted")
			metric.SetDescription("The total number of accepted client connections")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			metric.IntSum().SetIsMonotonic(true)
			metric.IntSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)

			return metric
		},
	},
	&metricImpl{
		"nginx.connections_active",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("nginx.connections_active")
			metric.SetDescription("The current number of open connections")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)

			return metric
		},
	},
	&metricImpl{
		"nginx.connections_handled",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("nginx.connections_handled")
			metric.SetDescription("The total number of handled connections. Generally, the parameter value is the same as nginx.connections_accepted unless some resource limits have been reached (for example, the worker_connections limit).")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			metric.IntSum().SetIsMonotonic(true)
			metric.IntSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)

			return metric
		},
	},
	&metricImpl{
		"nginx.connections_reading",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("nginx.connections_reading")
			metric.SetDescription("The current number of connections where nginx is reading the request headerhe current number of open connections")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)

			return metric
		},
	},
	&metricImpl{
		"nginx.connections_waiting",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("nginx.connections_waiting")
			metric.SetDescription("The current number of idle client connections waiting for a request.")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)

			return metric
		},
	},
	&metricImpl{
		"nginx.connections_writing",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("nginx.connections_writing")
			metric.SetDescription("The current number of connections where nginx is writing the response back to the client.")
			metric.SetUnit("connections")
			metric.SetDataType(pdata.MetricDataTypeIntGauge)

			return metric
		},
	},
	&metricImpl{
		"nginx.requests",
		func() pdata.Metric {
			metric := pdata.NewMetric()
			metric.InitEmpty()
			metric.SetName("nginx.requests")
			metric.SetDescription("Total number of requests made to the server since it started")
			metric.SetUnit("requests")
			metric.SetDataType(pdata.MetricDataTypeIntSum)
			metric.IntSum().SetIsMonotonic(true)
			metric.IntSum().SetAggregationTemporality(pdata.AggregationTemporalityCumulative)

			return metric
		},
	},
}

// M contains a set of methods for each metric that help with
// manipulating those metrics. M is an alias for Metrics
var M = Metrics

// Labels contains the possible metric labels that can be used.
var Labels = struct {
}{}

// L contains the possible metric labels that can be used. L is an alias for
// Labels.
var L = Labels