package api

import (
	"time"

	model "github.com/buyarella/shoppinglist_query/pkg/shoppinglist"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func ShoppingListFromAPIToModel(shoppingList *ShoppingList) *model.ShoppingList {
	entries := shoppingList.GetEntries()
	newEntries := make([]model.Entry, len(entries))
	for idx, entry := range entries {
		newEntries[idx] = ShoppingListEntryFromAPIToModel(entry)
	}

	users := shoppingList.GetSharedWith()
	sharedWith := make([]model.User, len(users))
	for idx, user := range users {
		sharedWith[idx] = UserFromAPIToModel(user)
	}

	return &model.ShoppingList{
		Name:        shoppingList.GetName(),
		DisplayName: shoppingList.GetDisplayName(),
		CreateTime:  TimestampToTime(shoppingList.GetCreateTime()),
		UpdateTime:  TimestampToTime(shoppingList.GetUpdateTime()),
		Entries:     newEntries,
		Creator:     UserFromAPIToModel(shoppingList.GetCreator()),
		SharedWith:  sharedWith,
	}
}

func ShoppingListEntryFromAPIToModel(entry *ShoppingListEntry) model.Entry {
	return model.Entry{
		Name:        entry.GetName(),
		CreateTime:  TimestampToTime(entry.GetCreateTime()),
		UpdateTime:  TimestampToTime(entry.GetUpdateTime()),
		BoughtTime:  TimestampToTime(entry.GetBoughtTime()),
		RemovedTime: TimestampToTime(entry.GetRemovedTime()),
		Item:        ItemFromAPIToModel(entry.GetItem()),
		Amount:      AmountFromAPIToModel(entry.GetAmount()),
		Requestor:   UserFromAPIToModel(entry.GetRequestor()),
		Offering:    OfferingFromAPIToModel(entry.GetOffering()),
	}
}

func ItemFromAPIToModel(item *Item) model.Item {
	return model.Item{
		Name:        item.GetName(),
		DisplayName: item.GetDisplayName(),
		ImageURL:    item.GetImageUrl(),
	}
}

func OfferingFromAPIToModel(offering *Offering) model.Offering {
	return model.Offering{
		Name:     offering.GetName(),
		Item:     ItemFromAPIToModel(offering.GetItem()),
		Vendor:   VendorFromAPIToModel(offering.GetVendor()),
		Price:    offering.GetPrice(),
		Currency: model.Currency(offering.GetCurrency()),
	}
}

func AmountFromAPIToModel(amount *Amount) model.Amount {
	return model.Amount{
		Count: amount.GetCount(),
		Unit:  model.Unit(amount.GetUnit()),
	}
}

func VendorFromAPIToModel(vendor *Vendor) model.Vendor {
	return model.Vendor{
		Name:        vendor.GetName(),
		DisplayName: vendor.GetDisplayName(),
		CreateTime:  TimestampToTime(vendor.GetCreateTime()),
		UpdateTime:  TimestampToTime(vendor.GetUpdateTime()),
	}
}

func UserFromAPIToModel(user *User) model.User {
	return model.User{
		Name:     user.GetName(),
		Email:    user.GetEmail(),
		Fistname: user.GetFistname(),
		Lastname: user.GetLastname(),
	}
}

func TimestampToTime(timestamp *timestamp.Timestamp) time.Time {
	return time.Unix(timestamp.GetSeconds(), int64(timestamp.GetNanos()))
}

func ShoppingListFromModelToAPI(shoppingList model.ShoppingList) *ShoppingList {
	entries := shoppingList.Entries
	allEntries := make([]*ShoppingListEntry, len(entries))
	for idx, entry := range entries {
		allEntries[idx] = ShoppingListEntryFromModelToAPI(entry)
	}

	users := shoppingList.SharedWith
	sharedWith := make([]*User, len(users))
	for idx, user := range users {
		sharedWith[idx] = UserFromModelToAPI(user)
	}

	return &ShoppingList{
		Name:        shoppingList.Name,
		DisplayName: shoppingList.DisplayName,
		CreateTime:  TimeToTimestamp(shoppingList.CreateTime),
		UpdateTime:  TimeToTimestamp(shoppingList.UpdateTime),
		Entries:     allEntries,
		Creator:     UserFromModelToAPI(shoppingList.Creator),
		SharedWith:  sharedWith,
	}
}

func ShoppingListEntryFromModelToAPI(entry model.Entry) *ShoppingListEntry {
	return &ShoppingListEntry{
		Name:        entry.Name,
		CreateTime:  TimeToTimestamp(entry.CreateTime),
		UpdateTime:  TimeToTimestamp(entry.UpdateTime),
		BoughtTime:  TimeToTimestamp(entry.BoughtTime),
		RemovedTime: TimeToTimestamp(entry.RemovedTime),
		Item:        ItemFromModelToAPI(entry.Item),
		Amount:      AmountFromModelToAPI(entry.Amount),
		Requestor:   UserFromModelToAPI(entry.Requestor),
		Offering:    OfferingFromModelToAPI(entry.Offering),
	}
}

func ItemFromModelToAPI(item model.Item) *Item {
	return &Item{
		Name:        item.Name,
		DisplayName: item.DisplayName,
		ImageUrl:    item.ImageURL,
	}
}

func OfferingFromModelToAPI(offering model.Offering) *Offering {
	return &Offering{
		Name:     offering.Name,
		Item:     ItemFromModelToAPI(offering.Item),
		Vendor:   VendorFromModelToAPI(offering.Vendor),
		Price:    offering.Price,
		Currency: Currency(offering.Currency),
	}
}

func AmountFromModelToAPI(amount model.Amount) *Amount {
	return &Amount{
		Count: amount.Count,
		Unit:  Unit(amount.Unit),
	}
}

func VendorFromModelToAPI(vendor model.Vendor) *Vendor {
	return &Vendor{
		Name:        vendor.Name,
		DisplayName: vendor.DisplayName,
		CreateTime:  TimeToTimestamp(vendor.CreateTime),
		UpdateTime:  TimeToTimestamp(vendor.UpdateTime),
	}
}

func UserFromModelToAPI(user model.User) *User {
	return &User{
		Name:     user.Name,
		Email:    user.Email,
		Fistname: user.Fistname,
		Lastname: user.Lastname,
	}
}

func TimeToTimestamp(tim time.Time) *timestamp.Timestamp {
	return &timestamp.Timestamp{Seconds: tim.Unix(), Nanos: int32(tim.Nanosecond())}
}
