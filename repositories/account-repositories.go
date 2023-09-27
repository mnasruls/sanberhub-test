package repositories

import (
	"log"
	"sanberhub-test/entities/models"
	"sanberhub-test/helpers"

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

	tx := acc.modelsDB.Debug().Create(account)
	if tx.Error != nil {
		return tx.Error
	}

	log.Println("create user account successfully")

	return nil
}

func (acc *AccountRepositories) GetAccount(accountNumber *string) (*models.Account, string, error) {
	log.Println("get user account data ...")

	var userAccount models.Account

	tx := acc.modelsDB.Debug().Where("account_number = ?", *accountNumber).First(&userAccount)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, helpers.NOT_FOUND, nil
	}

	log.Println("get user account data successfully")
	return &userAccount, "", nil
}

func (acc *AccountRepositories) UpdateBalance(account *models.Account) error {
	log.Println("update balance in process ...")

	tx := acc.modelsDB.Debug().Where("account_number = ?", account.AccountNumber).Updates(account)
	if tx.Error != nil {
		return tx.Error
	}

	log.Println("update balance successfully")
	return nil
}
