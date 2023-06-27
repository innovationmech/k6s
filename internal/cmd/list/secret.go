package list

import (
	"context"
	"fmt"
	"github.com/innovationmech/k6s/internal/config"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"log"
)

var listSecretsCmd = &cobra.Command{
	Use:   "secrets",
	Short: "List all secrets",
	Run: func(cmd *cobra.Command, args []string) {
		listSecrets()
	},
}

func listSecrets() {
	clientset, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Error getting config: %v", err)
	}

	secrets, err := clientset.CoreV1().Secrets("").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to get secrets: %v", err)
	}

	for _, secret := range secrets.Items {
		fmt.Printf("Secret Name: %s\n", secret.Name)
	}
}
