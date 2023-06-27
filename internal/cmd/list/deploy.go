package list

import (
	"context"
	"fmt"
	"github.com/innovationmech/k6s/internal/config"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

var listDeploymentsCmd = &cobra.Command{
	Use:   "deploy",
	Short: "List all deployments",
	Run: func(cmd *cobra.Command, args []string) {
		listDeployments()
	},
}

func listDeployments() {
	clientset, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	deployments, err := clientset.AppsV1().Deployments("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get deployments: %v", err)
	}

	for _, deployment := range deployments.Items {
		fmt.Printf("Deployment Name: %s\n", deployment.Name)
	}
}
