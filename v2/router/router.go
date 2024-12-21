package router

import (
	"fmt"
	"os"
	"strconv"
	"time"
	"yuth-optician-api/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *fiber.App {
	app := fiber.New()

	// arg group
	api := app.Group("/api")

	apiV1(api, db)

	return app
}

func apiV1(r fiber.Router, db *gorm.DB) {
	v1 := r.Group("/v1")
	v1.Use("/books", authRequired)
	
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Hello": "World"})
	})
	v1.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(models.GetAllBooks(db))
	})

	// BOOKS API

	// CREATE
	v1.Post("/books", func(c *fiber.Ctx) error {
		var book models.Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		err := models.CreateBook(db, &book)
		if err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "invalid json"})
		}
		return c.JSON(fiber.Map{
			"message": "CREATE SUCCESS",
		})
	})

	// READ
	v1.Get("/books/:id", func(c *fiber.Ctx) error {
		var book *models.Book
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}
		book = models.GetBook(db, id)
		if book.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "no books"})
		}
		return c.JSON(book)
	})

	// UPDATE
	// error : if id is not exists  will still create new
	v1.Put("books/:id", func(c *fiber.Ctx) error {
		var book models.Book
		if err := c.BodyParser(&book); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": err.Error()})
		}
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}
		book.ID = id
		err2 := models.UpdateBook(db, &book)
		if err2 != nil {
			return c.Status(404).JSON(fiber.Map{"error": "id not exists"})
		}
		return nil
	})

	// DELETE
	v1.Delete("books/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.SendStatus(fiber.ErrBadRequest.Code)
		}
		var book models.Book = *models.GetBook(db, id)
		if book.ID == 0 {
			return c.Status(400).JSON(fiber.Map{"error": "no id in db"})
		}
		err2 := models.DeleteBook(db, id)
		if err2 != nil {
			return c.Status(400).JSON(fiber.Map{"error": "no id in db"})
		}
		return nil
	})

	// USER API

	// REGISTER
	v1.Post("register", func(c *fiber.Ctx) error {
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
		}

		err := models.CreateUser(db, user)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "cant save user to db"})
		}
		return c.JSON(fiber.Map{"message": "CREATE USER SUCCESS"})
	})

	// LOGIN
	v1.Post("login", func(c *fiber.Ctx) error {
		user := new(models.User)
		if err := c.BodyParser(user); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "invalid json"})
		}

		token, err := models.LoginUser(db, user)
		if err != nil {
			return c.Status(401).JSON(fiber.Map{"error": "invalid credentials"})
		}

		// Set cookie
		c.Cookie(&fiber.Cookie{
			Name:     "jwt",
			Value:    token,
			Expires:  time.Now().Add(time.Hour * 72),
			HTTPOnly: true,
		})

		return c.JSON(fiber.Map{"message": "login successful!"})
	})
}

func apiV2(r fiber.Router, db *gorm.DB) {
	v2 := r.Group("/v2")
	v2.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"Hello": "World"})
	})
}


func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY") // do not leak this

	token, err := jwt.ParseWithClaims(cookie, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})	

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{"error": "unauthorized"})
	}
	claim := token.Claims.(jwt.MapClaims)
	fmt.Println(claim["user_id"])

	return c.Next()
}