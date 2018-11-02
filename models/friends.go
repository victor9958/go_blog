package models

import (
	_ "github.com/go-sql-driver/mysql"
)

type Friends struct {
	Id int
	FriendId int
	UserId int
	GroupId int
	//Group *Group `orm:"rel(one)"`
	RemarkName string `orm:"size(50)"`
}

