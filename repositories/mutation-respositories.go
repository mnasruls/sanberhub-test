package repositories

import (
	"log"
	"sanberhub-test/entities/models"
	"sanberhub-test/helpers"

	"gorm.io/gorm"
)

type MutationRepositories struct {
	modelsDB *gorm.DB
}

func NewMutationRepositories(gormDb *gorm.DB) *MutationRepositories {
	return &MutationRepositories{
		modelsDB: gormDb,
	}
}

func (mut *MutationRepositories) AddMutationHistory(mutation *models.Mutation) error {
	log.Println("add mutation history ...")

	tx := mut.modelsDB.Debug().Create(mutation)
	if tx.Error != nil {
		return tx.Error
	}

	log.Println("add mutation history successfully")

	return nil
}

func (mut *MutationRepositories) GetMutation(accountNumber *string) (*[]models.Mutation, string, error) {
	log.Println("get user mutation data ...")

	var userMutations []models.Mutation

	tx := mut.modelsDB.Debug().Where("user_account_number = ?", accountNumber).Find(&userMutations)
	if tx.Error != nil {
		return nil, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, helpers.NOT_FOUND, nil
	}

	log.Println("get user mutation data ...")

	return &userMutations, "", nil
}
