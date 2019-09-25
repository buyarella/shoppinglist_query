package api

import (
	"context"
	"time"

	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func New() *API {
	return &API{}
}

func (it *API) GetAllShoppingLists(ctx context.Context, request *GetAllShoppingListsRequest) (*GetAllShoppingListsResponse, error) {
	return &GetAllShoppingListsResponse{
		ItemsToBuy: []*ShoppingList{
			shoppingList,
		},
	}, nil
}

func (it *API) GetShoppingList(ctx context.Context, request *GetShoppingListRequest) (*ShoppingList, error) {
	return shoppingList, nil
}

func (it *API) GetAllItems(ctx context.Context, request *GetAllItemsRequest) (*GetAllItemsResponse, error) {
	return &GetAllItemsResponse{
		Items: items,
	}, nil
}

type API struct {
}

var shoppingList = &ShoppingList{
	Entries:     shoppinglistEntries,
	Name:        shoppingListName,
	DisplayName: "My First Shopping List",
}

var shoppinglistEntries = []*ShoppingListEntry{
	{
		Name:       shoppingListName + "banana",
		CreateTime: &timestamp.Timestamp{Seconds: time.Now().UTC().Unix()},
		Item:       items[0],
		Amount:     &Amount{Count: 5, Unit: Unit_PIECES},
		Offering: &Offering{
			Name:   vendorName + "/offerings/banana",
			Vendor: vendors[0],
			Price:  0.78,
			Per:    &Amount{Count: 1, Unit: Unit_PIECES},
		},
	},
	{
		Name:       shoppingListName + "bread",
		CreateTime: &timestamp.Timestamp{Seconds: time.Now().UTC().Unix()},
		Item:       items[3],
		Amount:     &Amount{Count: 1, Unit: Unit_PIECES},
		Offering: &Offering{
			Name:   vendorName + "/offerings/bread",
			Vendor: vendors[0],
			Price:  1.53,
			Per:    &Amount{Count: 1, Unit: Unit_PIECES},
		},
	},
	{
		Name:       shoppingListName + "strawberry",
		CreateTime: &timestamp.Timestamp{Seconds: time.Now().UTC().Unix()},
		Item:       items[3],
		Amount:     &Amount{Count: 1, Unit: Unit_KILO_GRAMM},
		Offering: &Offering{
			Name:   vendorName + "/offerings/strawberry",
			Vendor: vendors[0],
			Price:  2.5,
			Per:    &Amount{Count: 1, Unit: Unit_KILO_GRAMM},
		},
	},
}

var shoppingListName = "/typus/shoppinglists/123-23123-23523-123/"

var items = []*Item{
	{
		Name:        "buyerella/items/food/banana",
		DisplayName: "Banana",
		ImageUrl:    "https://cdn.mos.cms.futurecdn.net/42E9as7NaTaAi4A6JcuFwG-320-80.jpg",
	},
	{
		Name:        "buyerella/items/food/strawberry",
		DisplayName: "Strawberry",
		ImageUrl:    "https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/PerfectStrawberry.jpg/220px-PerfectStrawberry.jpg",
	},
	{
		Name:        "buyerella/items/food/kiwi",
		DisplayName: "Kiwi",
		ImageUrl:    "https://5.imimg.com/data5/KU/QG/MY-41533060/kiwi-fruit-500x500.jpg",
	},
	{
		Name:        "buyerella/items/food/bread",
		DisplayName: "Bread",
		ImageUrl:    "https://www.browneyedbaker.com/wp-content/uploads/2016/05/white-bread-51-600-600x400.jpg",
	},
	{
		Name:        "buyerella/items/hygiene/toilet-paper",
		DisplayName: "Toilet Paper",
		ImageUrl:    "https://www.staples-3p.com/s7/is/image/Staples/sp41688553_sc7?wid=512&hei=512",
	},
}

var vendorName = "vendors/my-store"

var vendors = []*Vendor{
	{
		Name:        vendorName,
		DisplayName: "My Store",
	},
}
