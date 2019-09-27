package shoppinglist

import (
	"time"

	linq "github.com/ahmetb/go-linq"

	// load postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repository interface {
	GetAllItems() []Item
	GetAllShoppingLists(user User) []ShoppingList
	GetShoppingList(shoppingList ShoppingList) (ShoppingList, error)
	GetActiveShoppingList() (ShoppingList, error)
}

func NewRepository(databaseConnectionString string) (Repository, error) {
	return &repository{
		shoppingLists: shoppingLists,
		items:         items,
		users:         make([]User, 0),
	}, nil
}

func (it *repository) GetAllItems() []Item {
	return it.items
}

func (it *repository) GetAllShoppingLists(user User) []ShoppingList {
	return it.shoppingLists
}

func (it *repository) GetShoppingList(shoppingList ShoppingList) (ShoppingList, error) {
	found := linq.From(it.shoppingLists).
		Where(func(sl interface{}) bool { return sl.(ShoppingList).Name == shoppingList.Name }).
		First()
	foundShoppingList, ok := found.(ShoppingList)
	if !ok {
		return ShoppingList{}, NotFound.New("no such shoppinglist: %v", shoppingList.Name)
	}
	return foundShoppingList, nil
}

func (it *repository) GetActiveShoppingList() (ShoppingList, error) {
	found := linq.From(it.shoppingLists).
		Where(func(sl interface{}) bool { return sl.(ShoppingList).IsActive }).
		First()
	foundShoppingList, ok := found.(ShoppingList)
	if !ok {
		return ShoppingList{}, NotFound.New("no active shoppinglist")
	}
	return foundShoppingList, nil
}

type repository struct {
	shoppingLists []ShoppingList
	items         []Item
	users         []User
}

var shoppingLists = []ShoppingList{
	{
		Entries:     shoppinglistEntries,
		Name:        shoppingListName,
		DisplayName: "My First Shopping List",
	},
	{
		Entries:     []Entry{shoppinglistEntries[0], shoppinglistEntries[1], shoppinglistEntries[3]},
		Name:        "/typus/shoppinglists/123456-456-789",
		DisplayName: "Another Shopping List",
	},
}

var shoppinglistEntries = []Entry{
	{
		Name:       shoppingListName + "banana",
		CreateTime: time.Now().UTC(),
		Item:       items[0],
		Amount:     Amount{Count: 5, Unit: Pieces},
		Offering: Offering{
			Name:   vendorName + "/offerings/banana",
			Vendor: vendors[0],
			Price:  0.78,
			Per:    Amount{Count: 1, Unit: Pieces},
		},
	},
	{
		Name:       shoppingListName + "bread",
		CreateTime: time.Now().UTC(),
		Item:       items[3],
		Amount:     Amount{Count: 1, Unit: Pieces},
		Offering: Offering{
			Name:   vendorName + "/offerings/bread",
			Vendor: vendors[0],
			Price:  1.53,
			Per:    Amount{Count: 1, Unit: Pieces},
		},
	},
	{
		Name:       shoppingListName + "strawberry",
		CreateTime: time.Now().UTC(),
		Item:       items[1],
		Amount:     Amount{Count: 1, Unit: KiloGramm},
		Offering: Offering{
			Name:   vendorName + "/offerings/strawberry",
			Vendor: vendors[0],
			Price:  2.5,
			Per:    Amount{Count: 1, Unit: KiloGramm},
		},
	},
	{
		Name:       shoppingListName + "kiwi",
		CreateTime: time.Now().UTC(),
		Item:       items[2],
		Amount:     Amount{Count: 1, Unit: KiloGramm},
		Offering: Offering{
			Name:   vendorName + "/offerings/kiwi",
			Vendor: vendors[0],
			Price:  2.5,
			Per:    Amount{Count: 1, Unit: KiloGramm},
		},
	},
	{
		Name:       shoppingListName + "toilet-paper",
		CreateTime: time.Now().UTC(),
		Item:       items[4],
		Amount:     Amount{Count: 1, Unit: KiloGramm},
		Offering: Offering{
			Name:   vendorName + "/offerings/toilet-paper",
			Vendor: vendors[0],
			Price:  3.5,
			Per:    Amount{Count: 1, Unit: KiloGramm},
		},
	},
}

var shoppingListName = "/typus/shoppinglists/123-23123-23523-123/"

var vendorName = "vendors/my-store"

var vendors = []Vendor{
	{
		Name:        vendorName,
		DisplayName: "My Store",
	},
}

var items = []Item{
	{
		Name:        "buyerella/items/food/banana",
		DisplayName: "Banana",
		ImageURL:    "https://cdn.mos.cms.futurecdn.net/42E9as7NaTaAi4A6JcuFwG-320-80.jpg",
	},
	{
		Name:        "buyerella/items/food/strawberry",
		DisplayName: "Strawberry",
		ImageURL:    "https://upload.wikimedia.org/wikipedia/commons/thumb/2/29/PerfectStrawberry.jpg/220px-PerfectStrawberry.jpg",
	},
	{
		Name:        "buyerella/items/food/kiwi",
		DisplayName: "Kiwi",
		ImageURL:    "https://5.imimg.com/data5/KU/QG/MY-41533060/kiwi-fruit-500x500.jpg",
	},
	{
		Name:        "buyerella/items/food/bread",
		DisplayName: "Bread",
		ImageURL:    "https://www.browneyedbaker.com/wp-content/uploads/2016/05/white-bread-51-600-600x400.jpg",
	},
	{
		Name:        "buyerella/items/hygiene/toilet-paper",
		DisplayName: "Toilet Paper",
		ImageURL:    "https://www.staples-3p.com/s7/is/image/Staples/sp41688553_sc7?wid=512&hei=512",
	},
}
