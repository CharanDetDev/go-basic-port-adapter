package handler

import (
	"encoding/json"
	"strconv"

	"github.com/CharanDetDev/go-basic-port-adapter/repository"
	"github.com/CharanDetDev/go-basic-port-adapter/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type (
	personHandler struct {
		PersonService service.PersonService
	}

	PersonHandler interface {
		GetPersonByID(c *fiber.Ctx) error
	}
)

func NewPersonHandler(personService service.PersonService) PersonHandler {
	return &personHandler{
		PersonService: personService,
	}
}

func (handler *personHandler) GetPersonByID(c *fiber.Ctx) error {

	personId, err := strconv.Atoi(c.Params("personId"))
	if err != nil {
		return Response(c, fiber.StatusBadRequest, "Invalid param")
	}

	var resPerson repository.PersonModel
	err = handler.PersonService.GetPersonByID(personId, &resPerson)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Response(c, fiber.StatusOK, "Not found")
		} else {
			return Response(c, fiber.StatusInternalServerError, "Internal server error")
		}
	}

	return Response(c, fiber.StatusOK, resPerson)
}

func Response(c *fiber.Ctx, httpCode int, data interface{}) (err error) {

	js, _ := json.Marshal(data)
	c.Write(js)
	c.Status(httpCode)

	return nil
}
