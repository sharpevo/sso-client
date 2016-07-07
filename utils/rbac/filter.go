package rbac

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"sso-client/utils/userinfo"
)

func HasRole(userInfo userinfo.UserInfo, role string) bool {
	roleName := GetRoleName(role)
	beego.Debug("HasRole:", roleName, "in", userInfo.Roles)
	return userInfo.Roles[GetRoleName(role)]
}

func GetRoleName(role string) (roleName string) {
	appName := beego.AppConfig.String("appname")
	roleName = fmt.Sprintf("%s-%s", "admin", appName)
	return roleName
}

func RoleCheck(ctx *context.Context, role string) {
	redirect := beego.URLFor("HomeController.Get")

	userInfo := userinfo.GetUserInfo(ctx.Request)
	if !HasRole(userInfo, role) {
		beego.Debug("CheckRole:", role, "not in", userInfo.Roles)
		ctx.Redirect(302, redirect)
		return
	}
}

func RoleAdminCheck(ctx *context.Context) {
	roleName := GetRoleName("admin")
	beego.Debug("CheckAdmin:", roleName)
	RoleCheck(ctx, roleName)
}

func Check(pattern string, position int, functions ...beego.FilterFunc) {
	for _, function := range functions {
		beego.InsertFilter(pattern, position, function)

	}
}

func AdminCheck(paths ...string) {
	for _, path := range paths {
		Check(path, beego.BeforeRouter, RoleAdminCheck)
	}
}
