package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	books = append(books, Book{ID: 1, Title: "Ai Keng Larb", Author: "FujiSan"})
	books = append(books, Book{ID: 2, Title: "LarbNongKhai", Author: "FujiSan"})

	app.Post("/login", login)

	app.Use(checkMiddleware)
	bookGroup := app.Group("/book")
	bookGroup.Use(checkMiddleware)
	bookGroup.Get("/", getBooks)
	bookGroup.Get("/:id", getBook)
	bookGroup.Post("/", createBook)
	bookGroup.Put("/:id", updateBook)
	bookGroup.Delete("/:id", deleteBook)

	app.Post("/upload", uploadFile)

	app.Get("/test-html", testHtml)

	app.Get("/config", getEnv)

	// app.Get("/hello", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!")
	// })

	app.Listen(":8080")
}

func uploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.SaveFile(file, "./uploads/"+file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("File uploaded successfully")
}

func testHtml(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
		"book":  books,
	})
}

func getEnv(c *fiber.Ctx) error {
	secret := os.Getenv("SECRET")

	if secret == "" {
		secret = "defaultsecret"
	}
	return c.JSON(fiber.Map{
		"SECRET": secret,
	})
}

func checkMiddleware(c *fiber.Ctx) error {
	start := time.Now()
	fmt.Printf("URL = %S", "Method = %S", "Time = %s\n", c.OriginalURL(), c.Method(), start)
	return c.Next()
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var member = User{
	Email:    "saklaw33@gmail.com",
	Password: "12345678",
}

func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if user.Email != member.Email || user.Password != member.Password {
		return fiber.ErrUnauthorized
	}
	return c.JSON(fiber.Map{
		"message": "Login Success",
	})
}
