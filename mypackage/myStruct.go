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
type Tag struct{
	Tag string `json:"tag"`
	Type int64 `json:"type"`
	Id int64 `json:"id"`
}
type Passage struct{
	Id  int64 `json:"id"`
	Blog string `json"blog"`
}
func (Passage) TableName()string{
	return "passage"
}