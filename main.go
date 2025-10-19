package main

import (
	"User/models"
	"User/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Car struct {
	Brand string  `json:"brand"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (r *Repository) Create(ctx *fiber.Ctx) error {
	car := &Car{}
	err := ctx.BodyParser(car)
	if err != nil {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "can't parse request",
		})
		return err
	}

	if car.Price == 0 {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "price can't be 0",
		})
		return nil
	}

	if car.Brand == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"brand can't be empty",
		})
		return nil
	}

	if car.Name == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message":"name can't be empty",
		})
		return nil
	}
	
	err = r.DB.Create(car).Error
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "can't create car",
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "car was created successfully",
	})
	return nil
}

func (r *Repository) DeleteByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	car := &models.Cars{}
	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID can't be empty",
		})
		return nil
	}
	err := r.DB.Delete(car, id).Error
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "can't delete this car",
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "car was deleted",
	})
	return nil

}

func (r *Repository) DeleteAll(ctx *fiber.Ctx) error {
	cars := &[]models.Cars{}
	err := r.DB.Where("1 = 1").Delete(cars).Error
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "can't delete cars",
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "cars was deleted",
	})
	return nil
}

func (r *Repository) ChangePrice(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	car := &models.Cars{}
	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID can't be empty",
		})
		return nil
	}

	type changer struct {
		Price float64 `json:"price"`
	}

	var request changer
	err := ctx.BodyParser(&request)
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "can't parse request",
		})
		return err
	}

	err = r.DB.Model(car).Where("id = ?", id).Update("price", request.Price).Error
	if err != nil {
		ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "can't change price",
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "price was update",
	})
	return nil
}

func (r *Repository) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	car := &models.Cars{}
	if id == "" {
		ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID can't be empty",
		})
		return nil
	}
	err := r.DB.Where("id = ?", id).First(car).Error
	if err != nil {
		ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "can't found car",
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "car was founded",
		"data":    car,
	})
	return nil
}

func (r *Repository) GetCars(ctx *fiber.Ctx) error {
	cars := &[]models.Cars{}
	err := r.DB.Find(cars).Error
	if err != nil {
		ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"message": "can't found car",
		})
		return err
	}
	ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "car was founded",
		"data":    cars,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create", r.Create)
	api.Delete("/delete/:id", r.DeleteByID)
	api.Delete("/deleteall", r.DeleteAll)
	api.Patch("/change/:id", r.ChangePrice)
	api.Get("/car/:id", r.GetByID)
	api.Get("/cars", r.GetCars)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal(err)
	}
	r := Repository{DB: db}
	err = models.MigrateCars(db)
	if err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
