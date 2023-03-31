package service

import (
	"fmt"

	"github.com/CharanDetDev/go-basic-port-adapter/repository"
	"gorm.io/gorm"
)

type (
	personService struct {
		PersonRepo repository.PersonRepo
	}
	PersonService interface {
		GetPersonByID(personId int, person *repository.PersonModel) error
	}
)

func NewPersonService(personRepo repository.PersonRepo) PersonService {
	return &personService{
		PersonRepo: personRepo,
	}
}

func (service *personService) GetPersonByID(personId int, person *repository.PersonModel) error {

	err := service.PersonRepo.GetPersonByID(personId, person)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Println("Get person not found")
			return gorm.ErrRecordNotFound
		}
		return err
	}

	return nil
}
