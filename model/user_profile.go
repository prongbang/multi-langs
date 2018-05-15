package model

type UserProfile struct {
	ID        int64  `json:"id" gorm:"primary_key"`
	UserID    int64  `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Gender    string `json:"gender"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}

func (userPProfile *UserProfile) TableName() string {
	return "user_profile"
}
