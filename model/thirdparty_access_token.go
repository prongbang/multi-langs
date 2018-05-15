package model

type ThirdPartyAccessToken struct {
	ID     int64  `json:"id" gorm:"primary_key"`
	Token  string `json:"token"`
	AppKey string `json:"app_key"`
}

func (token *ThirdPartyAccessToken) TableName() string {
	return "thirdparty_access_token"
}
