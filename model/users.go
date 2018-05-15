package model

type Users struct {
	ID       int64  `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"-"`
	Enabled  int    `json:"enabled" gorm:"default:1"`
}

func (user *Users) TableName() string {
	return "users"
}
