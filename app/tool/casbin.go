package tool

import (
	"Go_Practice/app/models"
	"Go_Practice/internal/database"
	"log"

	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

var Enforcer *casbin.Enforcer

type AdminWithPermissions struct {
	models.Admin
	Roles []RoleWithPermissions
}

type RoleWithPermissions struct {
	models.Role
	Permissions []string
}

// casbin 設定
func SetupCasbin() {
	// 使用DB 當police
	adapter, err := gormadapter.NewAdapterByDB(database.Db)
	if err != nil {
		log.Fatal("Carbins adapter error: ", err)
	}

	Enforcer, err = casbin.NewEnforcer("./env/casbin_rbac_with_root_model.conf", adapter)
	if err != nil {
		log.Fatal("Carbins init error: ", err)
	}
}

func GetPermissionsByRole(user string) []string {
	// Carbins GetPermissionsForUser此Api基於 rbac model 下只能以role去取得permissions (非user)
	pArray := Enforcer.GetPermissionsForUser(user)
	var permissions []string

	for _, v := range pArray {
		permissions = append(permissions, v[2])
	}
	return util.RemoveDuplicateElement(permissions)
}

func GetAdminWithPermissions(admin models.Admin) AdminWithPermissions {
	roles := []RoleWithPermissions{}
	for _, v := range admin.Roles {
		roles = append(
			roles,
			GetRoleWithPermissions(v),
		)
	}

	return AdminWithPermissions{
		Admin: admin,
		Roles: roles,
	}
}

func GetRoleWithPermissions(role models.Role) RoleWithPermissions {
	return RoleWithPermissions{
		Role:        role,
		Permissions: GetPermissionsByRole("role" + strconv.FormatUint(role.Id, 10)),
	}
}
