package cmd

import (
	"github.com/Project-Prismatica/prismatica_integration_nessus/subscribers"
	"github.com/lileio/lile"
	"github.com/lileio/pubsub"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "up runs both RPC and pubub subscribers",
	Run: func(cmd *cobra.Command, args []string) {
		go func() {
			logrus.Fatal(lile.Serve())

		}()

		pubsub.Subscribe(&subscribers.PrismaticaIntegrationNessusServiceSubscriber{})
	},
}

func init() {
	RootCmd.AddCommand(upCmd)
}
