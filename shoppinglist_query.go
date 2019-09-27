package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/buyarella/shoppinglist_query/pkg/api"
	"github.com/buyarella/shoppinglist_query/pkg/config"
	"github.com/buyarella/shoppinglist_query/pkg/repository"

	"github.com/sirupsen/logrus"
)

var mainLog = logrus.WithField("name", "main")

func main() {
	cfg := config.Get()

	address := fmt.Sprintf(":%d", cfg.Port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		mainLog.Fatalf("could not create listener on address: %s", address)
	}

	shoppingListsRepository, err := repository.NewRepository(cfg.DatabaseConnectionString)
	if err != nil {
		panic(err)
	}

	apiServer := api.New(shoppingListsRepository)

	grpcServer := grpc.NewServer()
	api.RegisterShoppingListQueriesServer(grpcServer, apiServer)

	mainLog.Infof("started grpc server on '%v'", address)
	if err := grpcServer.Serve(listener); err != nil {
		mainLog.Fatalf("failed to serve grpc: '%v'", err)
	}
}
