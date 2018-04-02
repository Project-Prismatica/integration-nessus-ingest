package prismatica_integration_nessus

import (
	"sync"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

var (
	cm     = &sync.Mutex{}
	Client PrismaticaIntegrationNessusClient
)

func GetPrismaticaIntegrationNessusClient() (PrismaticaIntegrationNessusClient, error) {
	cm.Lock()
	defer cm.Unlock()

	if Client != nil {
		return Client, nil
	}

	logrus.Info("Creating prismatica_integration_nessus gRPC client")
	conn, err := grpc.Dial("prismatica_integration_nessus:80", grpc.DialOption(grpc.WithInsecure()))
	if err != nil {
		return nil, err
	}

	cli := NewPrismaticaIntegrationNessusClient(conn)
	Client = cli
	return cli, nil
}
