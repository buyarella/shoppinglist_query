package repository

import (
	"time"

	"github.com/buyarella/shoppinglist_query/pkg/repository/model"
	"github.com/jinzhu/gorm"

	// load postgres driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Repository interface {
	GetAllItems() []model.Item
	GetAllShoppingLists(user model.User) []model.ShoppingList
	GetShoppingList(shoppingList model.ShoppingList) (model.ShoppingList, error)
	GetActiveShoppingList() (model.ShoppingList, error)
}

func NewRepository(databaseConnectionString string) (Repository, error) {
	db, err := gorm.Open("postgres", databaseConnectionString)

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(
		&model.ShoppingList{},
		&model.ShoppingListEntry{},
		&model.Item{},
		&model.Amount{},
		&model.Offering{},
		&model.Vendor{},
		&model.User{},
	)

	for _, item := range items {
		toCreate := item
		db.Create(&toCreate)
	}

	for _, shoppingList := range shoppingLists {
		toCreate := shoppingList
		db.Create(&toCreate)
	}

	return &repository{db: db}, nil
}

func (it *repository) GetAllItems() []model.Item {
	var items []model.Item
	it.db.Find(&items)
	return items
}

func (it *repository) GetAllShoppingLists(user model.User) []model.ShoppingList {
	var shoppingLists []model.ShoppingList
	it.db.Find(&shoppingLists)
	return shoppingLists
}

func (it *repository) GetShoppingList(shoppingList model.ShoppingList) (model.ShoppingList, error) {
	var list model.ShoppingList
	it.db.Where(&shoppingList).Find(&list)
	return list, nil
}

func (it *repository) GetActiveShoppingList() (model.ShoppingList, error) {
	var list model.ShoppingList
	return list, nil
}

type repository struct {
	db *gorm.DB
}

var shoppingLists = []model.ShoppingList{
	{
		Entries:     shoppinglistEntries,
		Name:        shoppingListName,
		DisplayName: "My First Shopping List",
	},
	{
		Entries:     []model.ShoppingListEntry{shoppinglistEntries[0], shoppinglistEntries[1], shoppinglistEntries[3]},
		Name:        "/typus/shoppinglists/123456-456-789",
		DisplayName: "Another Shopping List",
	},
}

var shoppinglistEntries = []model.ShoppingListEntry{
	{
		Name:       shoppingListName + "banana",
		CreateTime: time.Now().UTC(),
		Item:       items[0],
		Amount:     model.Amount{Count: 5, Unit: model.Pieces},
		Offering: model.Offering{
			Name:   vendorName + "/offerings/banana",
			Vendor: vendors[0],
			Price:  0.78,
			Per:    model.Amount{Count: 1, Unit: model.Pieces},
		},
	},
	{
		Name:       shoppingListName + "bread",
		CreateTime: time.Now().UTC(),
		Item:       items[3],
		Amount:     model.Amount{Count: 1, Unit: model.Pieces},
		Offering: model.Offering{
			Name:   vendorName + "/offerings/bread",
			Vendor: vendors[0],
			Price:  1.53,
			Per:    model.Amount{Count: 1, Unit: model.Pieces},
		},
	},
	{
		Name:       shoppingListName + "strawberry",
		CreateTime: time.Now().UTC(),
		Item:       items[1],
		Amount:     model.Amount{Count: 1, Unit: model.KiloGramm},
		Offering: model.Offering{
			Name:   vendorName + "/offerings/strawberry",
			Vendor: vendors[0],
			Price:  2.5,
			Per:    model.Amount{Count: 1, Unit: model.KiloGramm},
		},
	},
	{
		Name:       shoppingListName + "kiwi",
		CreateTime: time.Now().UTC(),
		Item:       items[2],
		Amount:     model.Amount{Count: 1, Unit: model.KiloGramm},
		Offering: model.Offering{
			Name:   vendorName + "/offerings/kiwi",
			Vendor: vendors[0],
			Price:  2.5,
			Per:    model.Amount{Count: 1, Unit: model.KiloGramm},
		},
	},
	{
		Name:       shoppingListName + "toilet-paper",
		CreateTime: time.Now().UTC(),
		Item:       items[4],
		Amount:     model.Amount{Count: 1, Unit: model.KiloGramm},
		Offering: model.Offering{
			Name:   vendorName + "/offerings/toilet-paper",
			Vendor: vendors[0],
			Price:  3.5,
			Per:    model.Amount{Count: 1, Unit: model.KiloGramm},
		},
	},
}

var shoppingListName = "/typus/shoppinglists/123-23123-23523-123/"

var vendorName = "vendors/my-store"

var vendors = []model.Vendor{
	{
		Name:        vendorName,
		DisplayName: "My Store",
	},
}

var items = []model.Item{
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
