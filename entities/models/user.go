package models

type User struct {
	Id        int    `json:"id,omitempty" gorm:"column:id;primary_key;auto_increment"`
	Name      string `json:"name" gorm:"column:name"`
	NIK       string `json:"nik" gorm:"column:nik"`
	NoHP      string `json:"no_hp" gorm:"column:no_hp"`
	CreatedAt string `json:"created_at" gorm:"column:created_at"`
}

func (User) TableName() string {
	return "user"
}
