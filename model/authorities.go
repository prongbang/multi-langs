package model

type Authorities struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	UserID    int64  `json:"user_id"`
	Authority string `json:"authority" gorm:"default:'ROLE_USER'"`
}

func (auth *Authorities) TableName() string {
	return "authorities"
}
