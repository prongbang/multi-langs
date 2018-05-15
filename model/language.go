package model

import (
	"time"
)

type Language struct {
	ID       int64     `json:"id" gorm:"primary_key"`
	Key      string    `json:"key"`
	Name     string    `json:"name"`
	Icon     string    `json:"icon"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
	CreateBy int64     `json:"create_by"`
	UpdateBy int64     `json:"update_by"`
	Status   int       `json:"status" gorm:"default:1"`
}

func (lang *Language) TableName() string {
	return "language"
}
