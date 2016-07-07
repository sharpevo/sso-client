package userinfo

import (
	"github.com/astaxie/beego"
	"net/http"
	"strings"
)

type UserInfo struct {
	Id    string
	Name  string
	Email string
	Roles map[string]bool
}

func GetUserInfo(req *http.Request) (userInfo UserInfo) {
	roles := req.Header.Get("Igenetech-User-Roles")
	roleList := strings.Split(roles, ",")
	beego.Debug("RoleList:", roleList)
	roleMap := make(map[string]bool)
	for _, role := range roleList {
		roleMap[role] = true
	}
	userInfo = UserInfo{
		Id:    req.Header.Get("Igenetech-User-Id"),
		Name:  req.Header.Get("igenetech-user-name"),
		Email: req.Header.Get("Igenetech-User-Email"),
		Roles: roleMap,
	}
	return userInfo

}
