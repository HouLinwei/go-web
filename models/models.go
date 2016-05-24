package main

import (
	"github.com/astaxie/beego/orm"
)


type User struct {
	Id int
	name string
	passwd string
}

