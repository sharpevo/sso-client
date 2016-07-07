package main

import (
	"github.com/astaxie/beego"
	_ "sso-client/routers"
)

func main() {
	beego.Run()
}
