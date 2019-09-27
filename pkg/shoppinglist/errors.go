package shoppinglist

import "github.com/joomcode/errorx"

var namespace = errorx.NewNamespace("buyarella")

var NotFound = errorx.NewType(namespace, "not_found", errorx.NotFound())
