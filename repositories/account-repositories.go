package repositories

import (
	"log"
	"sanberhub-test/entities/models"

	"gorm.io/gorm"
)

type AccountRepositories struct {
	modelsDB *gorm.DB
}

func NewAccountRepositories(gormDB *gorm.DB) *AccountRepositories {
	return &AccountRepositories{
		modelsDB: gormDB,
	}
}

func (acc *AccountRepositories) CreateAccount(account *models.Account) error {
	log.Println("create user account in process ...")

	tx := acc.modelsDB.Create(account)
	if tx.Error != nil {
		return tx.Error
	}

	log.Println("create user account successfully")

	return nil
}
