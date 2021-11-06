package mypackage

import "time"

type Cdetail struct {
	Title string `json:"name"`
	Id    int64 `json:"id"`
	Date  time.Time `json:"date"`
	Tag   string `json:"tag"`
	Type  int64 `json:"type"`
}
func (Cdetail) TableName()string{
	return "cdetail"
}
type Tag struct{
	Tag string `json:"tag"`
	Type int64 `json:"type"`
}