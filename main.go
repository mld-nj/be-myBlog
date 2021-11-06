package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mld-nj/my_blog_be/mypackage"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	r:=gin.Default()
	r.GET("/cardDetail",func(c *gin.Context) {
		var Cdetails []mypackage.Cdetail
		db.Find(&Cdetails)
		dJson,err:=json.Marshal(Cdetails)
		if err!=nil{
			fmt.Println("json化错误")
		}
		c.JSON(http.StatusOK,string(dJson))
	})
	r.GET("/tags",func(c *gin.Context) {
		var tags []mypackage.Tag
		db.Model(&mypackage.Cdetail{}).Select("tag","type").Find(&tags)
		dJson,err:=json.Marshal(tags)
		if err!=nil{
			fmt.Println("json化错误")
		}
		c.JSON(http.StatusOK,string(dJson))
	})
	r.Run()
}