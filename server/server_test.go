package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/Project-Prismatica/prismatica_integration_nessus"
	"github.com/lileio/lile"
)

var s = PrismaticaIntegrationNessusServer{}
var cli prismatica_integration_nessus.PrismaticaIntegrationNessusClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		prismatica_integration_nessus.RegisterPrismaticaIntegrationNessusServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = prismatica_integration_nessus.NewPrismaticaIntegrationNessusClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
