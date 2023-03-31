package main

import (
	"context"
	"fmt"
	"time"

	"github.com/CharanDetDev/go-basic-port-adapter/handler"
	"github.com/CharanDetDev/go-basic-port-adapter/repository"
	"github.com/CharanDetDev/go-basic-port-adapter/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

// Trace is Print SQL statement
func (sqlLog *SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sqlStatement, _ := fc()
	fmt.Println(" ***** Generate GORM SQL Statement ***** | ", sqlStatement)
}

func main() {

	DBConnect, err := gorm.Open(
		mysql.Open("root:#demo#MySQL@tcp(localhost:3307)/demoMySQL?parseTime=True&loc=Local"),
		&gorm.Config{
			Logger: &SqlLogger{},
			// DryRun: true,
		},
	)
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
