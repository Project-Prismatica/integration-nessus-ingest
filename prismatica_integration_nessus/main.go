package main

import (
	"github.com/Project-Prismatica/prismatica_integration_nessus"
	"github.com/Project-Prismatica/prismatica_integration_nessus/prismatica_integration_nessus/cmd"
	"github.com/Project-Prismatica/prismatica_integration_nessus/server"
	"github.com/lileio/lile"
	"github.com/lileio/lile/fromenv"
	"github.com/lileio/pubsub"
	"google.golang.org/grpc"
)

func main() {
	s := &server.PrismaticaIntegrationNessusServer{}

	lile.Name("prismatica_integration_nessus")
	lile.Server(func(g *grpc.Server) {
		prismatica_integration_nessus.RegisterPrismaticaIntegrationNessusServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
	})

	cmd.Execute()
}
