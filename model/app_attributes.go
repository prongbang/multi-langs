package model

import (
	"time"
)

type AppAttributes struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	LabelKey  string    `json:"label_key"`
	LabelVal  string    `json:"label_val"`
	CreateAt  time.Time `json:"create_at"`
	UpdateAt  time.Time `json:"update_at"`
	CreateBy  int64     `json:"create_by"`
	UpdateBy  int64     `json:"update_by"`
	Status    int       `json:"status" gorm:"default:1"`
	AppLangID int64     `json:"app_lang_id"`
}

func (lang *AppAttributes) TableName() string {
	return "app_attributes"
}
