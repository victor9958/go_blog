package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}


func (c *IndexController) Get() {
	//str:=strconv.Itoa(66666)
	//c.Ctx.WriteString(str)

	c.TplName = "index.html"
}