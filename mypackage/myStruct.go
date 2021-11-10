package mypackage

import "time"

type Cdetail struct {
	Title string `json:"name"`
	Id    int64 `json:"id"`
	Date  time.Time `json:"date"`
	Tag   string `json:"tag"`
	Type  int64 `json:"type"`
	Detail string `json:"detail"`
}
func (Cdetail) TableName()string{
	return "cdetail"
}
//获取所有tag
type Tag struct{
	Tag string `json:"tag"`
	Type int64 `json:"type"`
	Id int64 `json:"id"`
}
//获取文章
type Passage struct{
	Id  int64 `json:"id"`
	Blog string `json"blog"`
}
func (Passage) TableName()string{
	return "passage"
}
//获取文章数
type Num struct{
	Sum int64 `json:"sum"`
}
//获取tag种类
type Tagkind struct{
	// Tag string `json:"tagName"`
	Total     int64 `json:"total"`
}