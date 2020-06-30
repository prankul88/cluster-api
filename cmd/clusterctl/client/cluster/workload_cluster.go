/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cluster

import (
	"fmt"

	clusterv1 "sigs.k8s.io/cluster-api/api/v1alpha3"
)

type WorkloadCluster interface {
	//Get workload cluster kubeconfig
	GetKubeconfig(name string) error
}

// ProviderComponents implements ComponentsClient.
type providerComponent struct {
	proxy Proxy
}

func (p *providerComponent) GetKubeconfig(name string) error {
	// cs, err := p.proxy.NewClient()
	// if err != nil {
	// 	return err
	// }

	labels := map[string]string{
		clusterv1.ClusterLabelName: "workload",
	}

	namespace := "default"

	resources, err := p.proxy.ListResources(labels, namespace)
	if err != nil {
		fmt.Printf("error fetching")
	}

	fmt.Print(resources)

	return err
}

// newComponentsClient returns a providerComponents.
func newComponentClient(proxy Proxy) *providerComponent {
	return &providerComponent{
		proxy: proxy,
	}
}
