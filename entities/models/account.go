package models

type Account struct {
	Id            int    `json:"id,omitempty" gorm:"column:id;primary_key;auto_increment"`
	UserId        int    `json:"user_id" gorm:"column:user_id"`
	AccountNumber string `json:"account_number" gorm:"column:account_number"`
	Balance       int    `json:"balance" gorm:"column:balance"`
	CreatedAt     string `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     string `json:"updated_at" gorm:"column:updated_at"`
}

func (Account) TableName() string {
	return "account"
}
