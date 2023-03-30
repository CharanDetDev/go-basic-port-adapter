package repository

import (
	"gorm.io/gorm"
)

type (
	personRepo struct {
		DBConnect *gorm.DB
	}

	PersonRepo interface {
		GetPersonByID(personID int, person *PersonModel) error
	}

	PersonModel struct {
		PersonID  int    `gorm:"column:PersonID"`
		LastName  string `gorm:"column:LastName"`
		FirstName string `gorm:"column:FirstName"`
		Address   string `gorm:"column:Address"`
		City      string `gorm:"column:City"`
	}
)

func (m *PersonModel) TableName() string {
	return "Persons"
}

func NewPersonRepo(dbConnect *gorm.DB) PersonRepo {
	return &personRepo{
		DBConnect: dbConnect,
	}
}

func (repo *personRepo) GetPersonByID(personId int, person *PersonModel) error {
	err := repo.DBConnect.First(&person, personId).Error
	if err != nil {
		return err
	}
	return nil
}
