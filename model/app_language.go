package model

type AppLanguage struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	LangID int64  `json:"lang_id"`
	AppKey string `json:"app_key"`
}

func (lang *AppLanguage) TableName() string {
	return "app_language"
}
