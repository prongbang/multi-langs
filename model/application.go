package model

type Application struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"`
}

func (app *Application) TableName() string {
	return "application"
}
