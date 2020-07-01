/*
Copyright 2019 The Kubernetes Authors.

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

package client

import (
	"fmt"
)

//GetKubeconfigOptions carries all the options supported by GetKubeconfig
type GetKubeconfigOptions struct {
	// Kubeconfig defines the kubeconfig to use for accessing the management cluster. If empty,
	// default rules for kubeconfig discovery will be used.
	Kubeconfig Kubeconfig

	Name string

	//Out string
}

func (c *clusterctlClient) GetKubeconfig(options GetKubeconfigOptions) error {
	// gets access to the management cluster
	clusterClient, err := c.clusterClientFactory(options.Kubeconfig)
	if err != nil {
		fmt.Printf("Bc fir error")
	}

	//resources := []unstructured.Unstructured{}
	err = clusterClient.WorkloadCluster().GetKubeconfig(options.Name)

	return err
}
