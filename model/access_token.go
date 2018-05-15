package model

type AccessToken struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func (token *AccessToken) TableName() string {
	return "access_token"
}
