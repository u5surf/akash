package query

import (
	"github.com/ovrclk/akash/cmd/akash/session"
	"github.com/spf13/cobra"
)

func QueryCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "query [something]",
		Short: "Query something",
		Args:  cobra.ExactArgs(1),
	}

	session.AddFlagNode(cmd, cmd.PersistentFlags())

	cmd.AddCommand(
		queryAccountCommand(),
		queryDeploymentCommand(),
		queryDeploymentGroupCommand(),
		queryProviderCommand(),
		queryOrderCommand(),
		queryFulfillmentCommand(),
		queryLeaseCommand(),
	)

	return cmd
}
