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
	//获取blog的全部数据
	r.GET("/cardDetail",func(c *gin.Context) {
		var Cdetails []mypackage.Cdetail
		db.Find(&Cdetails)
		dJson,err:=json.Marshal(Cdetails)
		if err!=nil{
			fmt.Println("json化错误")
		}
		c.JSON(http.StatusOK,string(dJson))
	})
	//获取每个文章的标签
	r.GET("/tags",func(c *gin.Context) {
		var tags []mypackage.Tag
		db.Model(&mypackage.Cdetail{}).Select("tag","type").Find(&tags)
		dJson,err:=json.Marshal(tags)
		if err!=nil{
			fmt.Println("json化错误")
		}
		c.JSON(http.StatusOK,string(dJson))
	})
	//获取每个card对应的blog
	r.GET("/blog",func(c *gin.Context) {
		var passages []mypackage.Passage
		id:=c.DefaultQuery("id","1")
		db.Where("id=?",id).Find(&passages)
		c.String(http.StatusOK,string(passages[0].Blog))
	})
	//获取所有的文章数
	r.GET("/passageCounts",func(c *gin.Context) {
		var count int64
		db.Model(&mypackage.Cdetail{}).Distinct(`id`).Count(&count)
		num:=mypackage.Num{
			Sum: count,
		}
		dJson,err:=json.Marshal(num)
		if err!=nil{
			fmt.Println("json格式化错误")
		}
		c.JSON(http.StatusOK,string(dJson))
	})
	//获取所有文章中的tag种类数
	r.GET("/tagKinds",func(c *gin.Context) {
		var count int64
		db.Model(&mypackage.Cdetail{}).Distinct(`tag`).Count(&count)
		num:=mypackage.Tagkind{
			Total: count,
		}
		dJson,err:=json.Marshal(num)
		if err!=nil{
			fmt.Println("json格式化错误")
		}
		c.JSON(http.StatusOK,string(dJson))
	})
	//获取一共有多少种tag
	r.GET("/tagName",func(c *gin.Context) {
		var tagNames []mypackage.TagName
		db.Model(&mypackage.Cdetail{}).Select("tag","type").Group("tag").Find(&tagNames)
		dJson,err:=json.Marshal(tagNames)
		if err!=nil{
			fmt.Println("json格式化错误")
		}
		c.JSON(http.StatusOK,string(dJson))
	})
	r.Run()
}