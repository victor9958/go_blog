package main

import (
	_ "beego3/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beego3/models"
)

func insertCeshi(){
	o := orm.NewOrm()
	ceshiData := models.Blogger{}
	ceshiData.Name = "goCeshi"
	Id,err:=o.Insert(&ceshiData)
	if err!=nil {
		beego.Info("失败",err)
		return
	}
	beego.Info("成功",Id)
}

func insertBlog(){
	o := orm.NewOrm()
	ceshiData := models.Blog{}
	ceshiData.Title = "这是博客"
	ceshiData.Summary = "博客"
	ceshiData.Content = "这是博客的内容"
	Id,err:=o.Insert(&ceshiData)
	if err!=nil {
		beego.Info("失败",err)
		return
	}
	beego.Info("成功",Id)
}

func readCeshi(){
	o := orm.NewOrm()
	ceshiData := models.Blogger{Id:4825}
	err := o.Read(&ceshiData)
	if err!=nil {
		beego.Info("失败",err)
		return
	}
	beego.Info("成功",ceshiData)
}

func main() {
	insertBlog()
	//insertCeshi()
	//readCeshi()
	beego.Run()
}

