package list

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "List kubernetes resources",
	Long:  `List various types of resources in the Kubernetes cluster.`,
}

func init() {
	Cmd.AddCommand(listPodsCmd, listDeploymentsCmd, listServicesCmd, listSecretsCmd)
}
