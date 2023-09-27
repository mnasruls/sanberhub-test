package repositories

import (
	"log"
	"sanberhub-test/entities/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	modelsDB *gorm.DB
}

func NewUserRepository(gormDB *gorm.DB) *UserRepository {
	return &UserRepository{
		modelsDB: gormDB,
	}
}

func (usr *UserRepository) CreateUser(user *models.User) (int, error) {

	log.Println("create user in process ...")

	tx := usr.modelsDB.Debug().Create(user)
	if tx.Error != nil {
		return 0, tx.Error
	}

	log.Println("create user successfully")

	return user.Id, nil
}

func (usr *UserRepository) CheckUser(nik, noHp *string) (bool, error) {
	var users []models.User

	log.Println("checking for user ... ")

	tx := usr.modelsDB.Debug().Where("nik = ? OR no_hp = ?", *nik, *noHp).First(&users)
	if tx.Error != nil {
		return false, tx.Error
	}

	if tx.RowsAffected > 0 {
		return true, nil
	}

	return false, nil
}
