package main

// go run main.go
import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Person struct {
	Name    string `json:"name" form:"name"`
	YearOld int    `json:"year_old" form:"year_old"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000/",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	// middleware
	// Match any route - Phù hợp với bất kỳ tuyến đường nào
	app.Use(middlewareFirst)
	// Match all routes starting with /api - Khớp tất cả các tuyến đường bắt đầu bằng / api
	app.Use("/api", middlewareSecond)

	// routes
	app.Get("/adu/:name", handleName)
	app.Get("/adu/:name/:age", handleNameAge)
	app.Get("/api/adu", handleAdu)
	app.Get("/api/*", handleNotFound)
	app.Post("/person", handlePerson)

	app.Listen(":3000")
}

// middleware
func middlewareFirst(c *fiber.Ctx) error {
	fmt.Println("🥇 First middleware")
	return c.Next()
}
func middlewareSecond(c *fiber.Ctx) error {
	fmt.Println("🥈 Second middleware")
	return c.Next()
}

// routes
func handleName(c *fiber.Ctx) error {
	name := c.Params("name")
	return c.SendString(fmt.Sprintf("Hello, %s 👋!", name))
}
func handleNameAge(c *fiber.Ctx) error {
	name := c.Params("name")
	age := c.Params("age")
	return c.SendString(name + " : " + age + " years old")
}
func handleAdu(c *fiber.Ctx) error {
	return c.SendString("adu")
}
func handleNotFound(c *fiber.Ctx) error {
	msg := fmt.Sprintf("✋ %s not found", c.Params("*"))
	return c.SendString(msg)
}

func handlePerson(c *fiber.Ctx) error {
	person := &Person{}
	err := c.BodyParser(person)
	if err != nil {
		return c.SendString("BodyParser Error : " + err.Error())
	}
	return c.JSON(person)
}
