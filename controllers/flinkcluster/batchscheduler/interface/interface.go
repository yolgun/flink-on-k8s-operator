/*
Copyright 2020 Google LLC
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    https://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package schedulerinterface

import (
	"github.com/spotify/flink-on-k8s-operator/apis/flinkcluster/v1beta1"
	"github.com/spotify/flink-on-k8s-operator/controllers/flinkcluster/model"
)

// BatchScheduler is the general batch scheduler interface.
type BatchScheduler interface {
	// Name gets the name of the scheduler
	Name() string
	// Schedule reconciles batch scheduling
	Schedule(cluster *v1beta1.FlinkCluster, desired *model.DesiredClusterState) error
}
