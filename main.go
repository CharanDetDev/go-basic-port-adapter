package main

import (
	"fmt"

	"github.com/CharanDetDev/go-basic-port-adapter/handler"
	"github.com/CharanDetDev/go-basic-port-adapter/repository"
	"github.com/CharanDetDev/go-basic-port-adapter/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConnect *gorm.DB

func main() {

	DBConnect, err := gorm.Open(mysql.Open("root:#demo#MySQL@tcp(localhost:3307)/demoMySQL?parseTime=True&loc=Local"))
	if err != nil {
		fmt.Println("Connect DB Error | ", err.Error())
	}

	newPersonRepo := repository.NewPersonRepo(DBConnect)
	newPersonService := service.NewPersonService(newPersonRepo)
	newPersonHandler := handler.NewPersonHandler(newPersonService)

	app := fiber.New()
	app.Get("/person/:personId", newPersonHandler.GetPersonByID)

	app.Listen(":3000")

}
