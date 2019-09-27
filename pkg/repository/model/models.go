package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type ShoppingList struct {
	gorm.Model
	Name        string `gorm:"primary_key"`
	DisplayName string
	Entries     []ShoppingListEntry
	CreateTime  time.Time
	UpdateTime  time.Time
	Creator     User
	SharedWith  []User
}

type ShoppingListEntry struct {
	gorm.Model
	Name        string `gorm:"primary_key"`
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
	gorm.Model
	Name        string `gorm:"primary_key"`
	DisplayName string
	ImageURL    string
}

type Amount struct {
	gorm.Model
	Count int32
	Unit  Unit
}

type Offering struct {
	gorm.Model
	Name     string `gorm:"primary_key"`
	Item     Item
	Vendor   Vendor
	Price    float32
	Currency Currency
	Per      Amount
}

type Vendor struct {
	gorm.Model
	Name        string `gorm:"primary_key"`
	DisplayName string
	CreateTime  time.Time
	UpdateTime  time.Time
}

type User struct {
	gorm.Model
	Name     string `gorm:"primary_key"`
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
