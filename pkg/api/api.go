package api

import (
	"context"

	"github.com/buyarella/shoppinglist_query/pkg/repository"
	"github.com/buyarella/shoppinglist_query/pkg/repository/model"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func New(repo repository.Repository) *API {
	return &API{
		repository: repo,
	}
}

func (it *API) GetAllShoppingLists(ctx context.Context, request *GetAllShoppingListsRequest) (*GetAllShoppingListsResponse, error) {
	shoppingLists := it.repository.GetAllShoppingLists(model.User{})

	allShoppingLists := make([]*ShoppingList, len(shoppingLists))
	for _, list := range shoppingLists {
		allShoppingLists = append(allShoppingLists, ShoppingListFromModelToAPI(list))
	}

	return &GetAllShoppingListsResponse{
		ItemsToBuy: allShoppingLists,
	}, nil
}

func (it *API) GetShoppingList(ctx context.Context, request *GetShoppingListRequest) (*ShoppingList, error) {
	shoppingList, err := it.repository.GetShoppingList(model.ShoppingList{Name: request.GetName()})
	if err != nil {
		return nil, status.Error(codes.NotFound, "shoppingList was not found")
	}
	return ShoppingListFromModelToAPI(shoppingList), nil
}

func (it *API) GetActiveShoppingList(context.Context, *GetActiveShoppingListRequest) (*ShoppingList, error) {
	shoppingList, err := it.repository.GetActiveShoppingList()
	if err != nil {
		return nil, status.Error(codes.NotFound, "shoppingList was not found")
	}
	return ShoppingListFromModelToAPI(shoppingList), nil
}

func (it *API) GetAllItems(ctx context.Context, request *GetAllItemsRequest) (*GetAllItemsResponse, error) {
	items := it.repository.GetAllItems()

	allItems := make([]*Item, len(items))
	for _, item := range items {
		allItems = append(allItems, ItemFromModelToAPI(item))
	}

	return &GetAllItemsResponse{
		Items: allItems,
	}, nil
}

type API struct {
	repository repository.Repository
}
