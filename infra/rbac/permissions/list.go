package permissions

import "github.com/mikespook/gorbac"

// Define permissions
var (
	// !Order
	OrderCreate = gorbac.NewStdPermission("ORDER_CREATE")
	OrderCancel = gorbac.NewStdPermission("ORDER_CANCEL")
	OrderDelete = gorbac.NewStdPermission("ORDER_DELETE")
	OrderGet    = gorbac.NewStdPermission("ORDER_GET")
	OrderUpdate = gorbac.NewStdPermission("ORDER_UPDATE")
	// ! Product
	ProductCreate = gorbac.NewStdPermission("PRODUCT_CREATE")
	ProductDelete = gorbac.NewStdPermission("PRODUCT_DELETE")
	ProductGet    = gorbac.NewStdPermission("PRODUCT_GET")
	ProductUpdate = gorbac.NewStdPermission("PRODUCT_UPDATE")
	// ! user
	UserGet    = gorbac.NewStdPermission("UserGet")
	UserDelete = gorbac.NewStdPermission("UserDelete")
	UserUpdate = gorbac.NewStdPermission("UserUpdate")
)
