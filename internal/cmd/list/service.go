package list

import (
	"context"
	"fmt"
	"github.com/innovationmech/k6s/internal/config"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

var listServicesCmd = &cobra.Command{
	Use:   "services",
	Short: "List all services",
	Run: func(cmd *cobra.Command, args []string) {
		listServices()
	},
}

func listServices() {
	clientset, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	services, err := clientset.CoreV1().Services("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get services: %v", err)
	}

	for _, service := range services.Items {
		fmt.Printf("Service Name: %s\n", service.Name)
	}
}
