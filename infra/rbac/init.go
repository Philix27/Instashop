package rbac

import (
	"instashop/infra/rbac/permissions"
	"instashop/infra/rbac/roles"

	"github.com/mikespook/gorbac"
)

// Initialize RBAC with roles and permissions
func InitRBAC() *gorbac.RBAC {
	rbac := gorbac.New()

	// Define roles
	adminRole := gorbac.NewStdRole(roles.Admin)
	userRole := gorbac.NewStdRole(roles.User)

	//! Assign permissions to admin role
	adminRole.Assign(permissions.ProductDelete)
	adminRole.Assign(permissions.ProductUpdate)
	adminRole.Assign(permissions.ProductCreate)
	adminRole.Assign(permissions.ProductGet)
	// *order
	adminRole.Assign(permissions.OrderGet)
	// *user
	adminRole.Assign(permissions.UserGet)
	adminRole.Assign(permissions.UserUpdate)
	adminRole.Assign(permissions.UserDelete)

	//! Assign permissions to user role
	userRole.Assign(permissions.ProductGet)
	// *order
	userRole.Assign(permissions.OrderCreate)
	userRole.Assign(permissions.OrderDelete)
	userRole.Assign(permissions.OrderCancel)
	userRole.Assign(permissions.OrderGet)
	userRole.Assign(permissions.OrderUpdate)
	// *user
	userRole.Assign(permissions.UserGet)
	userRole.Assign(permissions.UserUpdate)
	userRole.Assign(permissions.UserDelete)

	// Add roles to RBAC
	rbac.Add(adminRole)
	rbac.Add(userRole)

	return rbac
}
