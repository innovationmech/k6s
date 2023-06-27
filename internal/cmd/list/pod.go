package list

import (
	"context"
	"fmt"
	"log"

	"github.com/innovationmech/k6s/internal/config"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var listPodsCmd = &cobra.Command{
	Use:   "pods",
	Short: "List all pods",
	Run: func(cmd *cobra.Command, args []string) {
		listPods()
	},
}

func listPods() {
	clientset, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	pods, err := clientset.CoreV1().Pods("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get pods: %v", err)
	}

	for _, pod := range pods.Items {
		fmt.Printf("Pod Name: %s\n", pod.Name)
	}
}
