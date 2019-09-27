package shoppinglist

import (
	"time"
)

type ShoppingList struct {
	Name        string
	DisplayName string
	Entries     []Entry
	CreateTime  time.Time
	UpdateTime  time.Time
	Creator     User
	SharedWith  []User
	IsActive    bool
}

type Entry struct {
	Name        string
	CreateTime  time.Time
	UpdateTime  time.Time
	BoughtTime  time.Time
	RemovedTime time.Time
	Item        Item
	Amount      Amount
	Requestor   User
	Offering    Offering
}

type Item struct {
	Name        string
	DisplayName string
	ImageURL    string
}

type Amount struct {
	Count int32
	Unit  Unit
}

type Offering struct {
	Name     string
	Item     Item
	Vendor   Vendor
	Price    float32
	Currency Currency
	Per      Amount
}

type Vendor struct {
	Name        string
	DisplayName string
	CreateTime  time.Time
	UpdateTime  time.Time
}

type User struct {
	Name     string
	Email    string
	Fistname string
	Lastname string
}

type Unit int32

const (
	Pieces    Unit = 0
	KiloGramm Unit = 1
	Gramm     Unit = 2
	Meter     Unit = 3
)

type Currency int32

const (
	Euro     Currency = 0
	Pount    Currency = 1
	UsDollar Currency = 2
)
