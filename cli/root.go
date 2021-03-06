// Copyright 2016 Google Inc.
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
///////////////////////////////////////////////////////////////////////////
package cli

import (
	"deploy"
	"log"
	"os"

	"github.com/spf13/cobra"

	"k8s.io/client-go/1.5/kubernetes"
	api "k8s.io/client-go/1.5/pkg/api/v1"
	"k8s.io/client-go/1.5/tools/clientcmd"
)

var (
	namespace string
	clientset *kubernetes.Clientset
	cfg       deploy.Service
)

func check(msg string, err error) {
	if err != nil {
		log.Println(msg, err)
		os.Exit(-1)
	}
}

// RootCmd for CLI
var RootCmd = &cobra.Command{
	Use:   "espcli",
	Short: "ESP deployment manager for Kubernetes",
	Long:  "A tool to deploy and monitor Extensible Service Proxy on a Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Please run 'espcli help' to see the list of available commands.")
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			clientcmd.NewDefaultClientConfigLoadingRules(),
			&clientcmd.ConfigOverrides{},
		).ClientConfig()
		if err != nil {
			log.Println("Cannot find the default Kubernetes configuration.")
			log.Println("Please check with kubectl your cluster config.")
			os.Exit(-2)
		}

		clientset, err = kubernetes.NewForConfig(config)
		if err != nil {
			log.Println("Cannot connect to the Kubernetes API: ", err)
			os.Exit(-2)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().StringVar(&namespace,
		"namespace", api.NamespaceDefault, "Specify Kubernetes namespace")
}
