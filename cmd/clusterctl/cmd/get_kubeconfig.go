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

package cmd

import (
	"github.com/spf13/cobra"
	"sigs.k8s.io/cluster-api/cmd/clusterctl/client"
)

type getKubeconfigOptions struct {
	name string
	//out  string
}

var gk = &getKubeconfigOptions{}

var getKubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Get kubeconfig for a workload cluster",
	Long: LongDesc(`
		Get kubeconfig for a workload cluster`),

	Example: Examples(`
		# Merge the workload cluster kubeconfig with management cluster kubeconfig
		clusterctl get kubeconfig --name my-workload-cluster

		# Output kubeconfig for a workload cluster in the output folder
		clusterctl get kubeconfig --name my-workload-cluster --out <dir> `),

	Args: cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGetKubeconfig(cmd)
	},
}

func init() {
	getCmd.Flags().StringVarP(&gk.name, "name", "n", "",
		"Name of the workload cluster to get kubeconfig")
	getCmd.MarkFlagRequired("name")
	//getCmd.Flags().StringVar(&gk.out, "out", "",
	//	"Merge managemnt cluster kubeconfig with manaement cluster kubeconfig")
	getCmd.AddCommand(getKubeconfigCmd)

}

func runGetKubeconfig(cmd *cobra.Command) error {
	c, err := client.New(cfgFile)
	if err != nil {
		return err
	}

	options := client.GetKubeconfigOptions{
		Kubeconfig: client.Kubeconfig{Path: initOpts.kubeconfig, Context: initOpts.kubeconfigContext},
		Name:       gk.name,
		//Out:         gk.out,
	}

	// if gk.name == "" {
	// 	return errors.Errorf("The name field is mandotory.")
	// }

	err = c.GetKubeconfig(options)

	return err

}
