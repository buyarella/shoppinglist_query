package api

import (
	"context"

	"github.com/buyarella/shoppinglist_query/pkg/shoppinglist"
	"github.com/joomcode/errorx"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func New(repo shoppinglist.Repository) *API {
	return &API{
		repository: repo,
	}
}

func (it *API) GetAllShoppingLists(ctx context.Context, request *GetAllShoppingListsRequest) (*GetAllShoppingListsResponse, error) {
	shoppingLists := it.repository.GetAllShoppingLists(shoppinglist.User{})

	allShoppingLists := make([]*ShoppingList, len(shoppingLists))
	for idx, list := range shoppingLists {
		allShoppingLists[idx] = ShoppingListFromModelToAPI(list)
	}

	return &GetAllShoppingListsResponse{
		ShoppingLists: allShoppingLists,
	}, nil
}

func (it *API) GetShoppingList(ctx context.Context, request *GetShoppingListRequest) (*ShoppingList, error) {
	shoppingList, err := it.repository.GetShoppingList(shoppinglist.ShoppingList{Name: request.GetName()})
	if err != nil {
		if errorx.HasTrait(err, errorx.NotFound()) {
			return nil, status.Error(codes.NotFound, request.GetName())
		}
		return nil, status.Errorf(codes.Internal, "could not retrieve shoppinglist: %v", request.GetName())
	}
	return ShoppingListFromModelToAPI(shoppingList), nil
}

func (it *API) GetActiveShoppingList(context.Context, *GetActiveShoppingListRequest) (*ShoppingList, error) {
	shoppingList, err := it.repository.GetActiveShoppingList()
	if err != nil {
		if errorx.HasTrait(err, errorx.NotFound()) {
			return nil, status.Error(codes.NotFound, "active shoppinglist")
		}
		return nil, status.Error(codes.Internal, "could not retrieve active shoppinglist")
	}
	return ShoppingListFromModelToAPI(shoppingList), nil
}

func (it *API) GetAllItems(ctx context.Context, request *GetAllItemsRequest) (*GetAllItemsResponse, error) {
	items := it.repository.GetAllItems()

	allItems := make([]*Item, len(items))
	for idx, item := range items {
		allItems[idx] = ItemFromModelToAPI(item)
	}

	return &GetAllItemsResponse{
		Items: allItems,
	}, nil
}

type API struct {
	repository shoppinglist.Repository
}
