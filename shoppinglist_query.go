package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/buyarella/shoppinglist_query/pkg/api"
	"github.com/buyarella/shoppinglist_query/pkg/config"
	"github.com/buyarella/shoppinglist_query/pkg/shoppinglist"
	"github.com/typusomega/poligo/pkg/policy"

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

	shoppingListsRepository, err := policy.HandleAll().
		Retry(
			policy.WithDurations(time.Second, time.Second, time.Second),
			policy.WithCallback(func(err error, retryCount int) { mainLog.WithError(err).Warn("retrying to connect to database") }),
		).
		Execute(context.Background(), func() (interface{}, error) { return shoppinglist.NewRepository(cfg.DatabaseConnectionString) })

	if err != nil {
		panic(err)
	}

	apiServer := api.New(shoppingListsRepository.(shoppinglist.Repository))

	resp, err := apiServer.GetAllShoppingLists(context.Background(), &api.GetAllShoppingListsRequest{})
	fmt.Print(resp)

	grpcServer := grpc.NewServer()
	api.RegisterShoppingListQueriesServer(grpcServer, apiServer)

	mainLog.Infof("started grpc server on '%v'", address)
	if err := grpcServer.Serve(listener); err != nil {
		mainLog.Fatalf("failed to serve grpc: '%v'", err)
	}
}
