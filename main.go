package main

import (
	_ "beego3/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beego3/models"
	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/gomodule/redigo/redis"
	"github.com/astaxie/beego/cache"
	"time"
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
	//server := ""


	//insertBlog()
	//insertCeshi()
	//readCeshi()
	cacheconn := beego.AppConfig.String("redis.conn")
	cachepwd := beego.AppConfig.String("redis.pwd")
	cachedb := beego.AppConfig.String("redis.cachedbname")
	cachestr := `{"key":"victor","conn":"`+cacheconn+`","dbNum":"`+cachedb+`","password":"`+ cachepwd+`"}`
	beego.Info(cachestr)
	bm,err :=cache.NewCache("redis",cachestr)
	if err != nil{
		beego.Info("redis作为缓存错误是:",err)
	}
	beego.Info(bm)
	bm.Put("name",1,10*time.Second)
	bm.Put("age",1000,1000*time.Second)
	beego.Info(bm.Get("name"))
	beego.Run()
}

