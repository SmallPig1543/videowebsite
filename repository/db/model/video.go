package model

type Video struct {
	ID       uint   `gorm:"primarykey" json:"ID"`
	Uid      uint   `json:"uid"` //发布者id
	Title    string `json:"title"`
	Types    string `json:"types"` //视频的种类
	Key      string //储存在oss中的key
	CreateAt string `json:"create_at"`
}
