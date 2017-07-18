// Copyright 2017 The Kubernetes Dashboard Authors.
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

package allobjects

import (
	metricapi "github.com/kubernetes/dashboard/src/app/backend/integration/metric/api"
	"github.com/kubernetes/dashboard/src/app/backend/resource/common"
	"github.com/kubernetes/dashboard/src/app/backend/resource/config"
	"github.com/kubernetes/dashboard/src/app/backend/resource/configmap"
	"github.com/kubernetes/dashboard/src/app/backend/resource/daemonset"
	"github.com/kubernetes/dashboard/src/app/backend/resource/dataselect"
	"github.com/kubernetes/dashboard/src/app/backend/resource/deployment"
	"github.com/kubernetes/dashboard/src/app/backend/resource/discovery"
	"github.com/kubernetes/dashboard/src/app/backend/resource/ingress"
	"github.com/kubernetes/dashboard/src/app/backend/resource/job"
	pvc "github.com/kubernetes/dashboard/src/app/backend/resource/persistentvolumeclaim"
	"github.com/kubernetes/dashboard/src/app/backend/resource/pod"
	"github.com/kubernetes/dashboard/src/app/backend/resource/replicaset"
	rc "github.com/kubernetes/dashboard/src/app/backend/resource/replicationcontroller"
	"github.com/kubernetes/dashboard/src/app/backend/resource/secret"
	"github.com/kubernetes/dashboard/src/app/backend/resource/service"
	"github.com/kubernetes/dashboard/src/app/backend/resource/statefulset"
	"github.com/kubernetes/dashboard/src/app/backend/resource/workload"
	"k8s.io/client-go/kubernetes"
)

// ObjectResult is a list of objects present in the given namespace
type ObjectResult struct {
	// Config and storage.
	ConfigMapList             configmap.ConfigMapList       `json:"configMapList"`
	PersistentVolumeClaimList pvc.PersistentVolumeClaimList `json:"persistentVolumeClaimList"`
	SecretList                secret.SecretList             `json:"secretList"`

	// Discovery and load balancing.
	ServiceList service.ServiceList `json:"serviceList"`
	IngressList ingress.IngressList `json:"ingressList"`

	// Workloads.
	DeploymentList            deployment.DeploymentList    `json:"deploymentList"`
	ReplicaSetList            replicaset.ReplicaSetList    `json:"replicaSetList"`
	JobList                   job.JobList                  `json:"jobList"`
	ReplicationControllerList rc.ReplicationControllerList `json:"replicationControllerList"`
	PodList                   pod.PodList                  `json:"podList"`
	DaemonSetList             daemonset.DaemonSetList      `json:"daemonSetList"`
	StatefulSetList           statefulset.StatefulSetList  `json:"statefulSetList"`

	// TODO(maciaszczykm): Third party resources.
}

func GetAllObjects(client *kubernetes.Clientset, metricClient metricapi.MetricClient,
	nsQuery *common.NamespaceQuery,
	dsQuery *dataselect.DataSelectQuery) (*ObjectResult, error) {

	configResources, err := config.GetConfig(client, nsQuery, dsQuery)
	if err != nil {
		return &ObjectResult{}, err
	}

	discoveryResources, err := discovery.GetDiscovery(client, nsQuery, dsQuery)
	if err != nil {
		return &ObjectResult{}, err
	}

	workloadsResources, err := workload.GetWorkloads(client, metricClient, nsQuery, dsQuery)
	if err != nil {
		return &ObjectResult{}, err
	}

	return &ObjectResult{
		// Config and storage.
		ConfigMapList:             configResources.ConfigMapList,
		PersistentVolumeClaimList: configResources.PersistentVolumeClaimList,
		SecretList:                configResources.SecretList,

		// Discovery and load balancing.
		ServiceList: discoveryResources.ServiceList,
		IngressList: discoveryResources.IngressList,

		// Workloads.
		DeploymentList:            workloadsResources.DeploymentList,
		ReplicaSetList:            workloadsResources.ReplicaSetList,
		JobList:                   workloadsResources.JobList,
		ReplicationControllerList: workloadsResources.ReplicationControllerList,
		PodList:                   workloadsResources.PodList,
		DaemonSetList:             workloadsResources.DaemonSetList,
		StatefulSetList:           workloadsResources.StatefulSetList,
	}, nil
}
