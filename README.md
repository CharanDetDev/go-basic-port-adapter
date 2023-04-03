# The Go Language Programming : Basic Port & Adapter

> ## **Package**
>
> _Fiber, GORM, MySQL_

- github.com/gofiber/fiber/v2
- gorm.io/gorm
- gorm.io/driver/mysql
- gorm.io/gorm/logger
- github.com/go-sql-driver/mysql

> ## **Example**

### package main

```golang

func main() {

        //* Database Connection
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

	newPersonRepo := repository.NewPersonRepo(DBConnect) // Inject Database Connection
	newPersonService := service.NewPersonService(newPersonRepo) // Inject Person Repository
	newPersonHandler := handler.NewPersonHandler(newPersonService) // Inject Person Service

	app := fiber.New()
	app.Get("/person/:personId", newPersonHandler.GetPersonByID) // เรียกใช้ Method GetPersonByID

	app.Listen(":3000")

}

```

### package handler

```golang

type (
	//* Adapter
	personHandler struct {
		PersonService service.PersonService
	}

        //* Interface Port
	PersonHandler interface {
		GetPersonByID(c *fiber.Ctx) error
	}
)

//* Function ที่ทำให้ Port และ Adapter เชื่อมต่อกัน
//* สำหรับ Web Framework เช่น  Fiber ทำการเรียกใช้งาน โดยจะต้อง Inject person Service มาด้วย
func NewPersonHandler(personService service.PersonService) PersonHandler {
	return &personHandler{
		PersonService: personService,
	}
}

//* Implementation Interface
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

```

### package service

```golang

type (
        //* Adapter
	personService struct {
		PersonRepo repository.PersonRepo
	}

        //* Interface Port
	PersonService interface {
		GetPersonByID(personId int, person *repository.PersonModel) error
	}
)

//* Function ที่ทำให้ Port และ Adapter เชื่อมต่อกัน
//* สำหรับ REST HTPP Handler ที่มี Dependency กับ person Service มาเรียกใช้โดยจะต้อง Inject person Repository มาด้วย
func NewPersonService(personRepo repository.PersonRepo) PersonService {
	return &personService{
		PersonRepo: personRepo,
	}
}

//* Implementation Interface
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

```

### package Repository

```golang

type (
        //* Adapter
	personRepo struct {
		DBConnect *gorm.DB
	}

        //* Interface Port
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

//* Function ที่ทำให้ Port และ Adapter เชื่อมต่อกัน
//* สสำหรับ Business-side/Domain ที่มี Dependency กับ person Repository มาเรียกใช้โดยจะต้อง Inject Database Connection มาด้วย 
func NewPersonRepo(dbConnect *gorm.DB) PersonRepo {
	return &personRepo{
		DBConnect: dbConnect,
	}
}

//* Implementation Interface
func (repo *personRepo) GetPersonByID(personId int, person *PersonModel) error {

	err := repo.DBConnect.First(&person, personId).Error
	if err != nil {
		return err
	}

	return nil
}

```
