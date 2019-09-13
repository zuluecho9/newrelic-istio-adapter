// Copyright 2019 New Relic Corporation
// Copyright 2018 Istio Authors
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

package trace

import (
	"github.com/newrelic/newrelic-istio-adapter/internal/nrsdk/telemetry"
	"github.com/newrelic/newrelic-istio-adapter/config"
	"istio.io/istio/mixer/pkg/adapter"
)

// BuildHandler returns a trace Handler with valid configuration.
func BuildHandler(_ *config.Params, h *telemetry.Harvester, env adapter.Env) (*Handler, error) {
	traceHandler := &Handler{
		logger:    env.Logger(),
		harvester: h,
	}
	return traceHandler, nil
}