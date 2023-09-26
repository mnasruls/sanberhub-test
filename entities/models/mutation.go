package models

type Mustation struct {
	Id              int    `json:"id,omitempty" gorm:"column:id;primary_key;auto_increment"`
	UserId          int    `json:"user_id" gorm:"column:user_id"`
	AccountNumber   int    `json:"account_number" gorm:"column:user_account_number"`
	KodeTransaksi   string `json:"transaction_code" gorm:"column:transaction_code"`
	TransactionTime string `json:"transaction_time" gorm:"column:transaction_time"`
	Nominal         int    `json:"nominal" gorm:"column:nominal"`
}
