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
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/myblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!=nil{
		panic(err)
	}
	r:=gin.Default()
	//跨域
	r.Use(Cors())
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
	r.GET("/blog",func(c *gin.Context) {
		var passages []mypackage.Passage
		id:=c.DefaultQuery("id","1")
		fmt.Println(id)
		db.Where("id=?",id).Find(&passages)
		// dJson,err:=json.Marshal(passages)
		// if err!=nil{
		// 	fmt.Println("json化错误")
		// }
		c.String(http.StatusOK,string(passages[0].Blog))
	})
	// r.GET("/blog",func(c *gin.Context) {
	// 	var passage mypackage.Passage
	// 	db.Where("id = ?",1).Find(passage)
	// 	dJson,err:=json.Marshal(passage)
	// 	if err!=nil{
	// 		fmt.Println("json化错误")
	// 	}
	// 	c.JSON(http.StatusOK,string(dJson))
	// })
	r.Run()
}