package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Todo struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {

	todos := []Todo{}

	fmt.Println("Hello world")

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("Ok")
	})

	app.Get("api/todos", func(c *fiber.Ctx) error {
		return c.JSON(todos)
	})

	app.Patch("api/todos/:id/done", func(c *fiber.Ctx) error {
		todoItem := &Todo{}
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(401).SendString("Invalid Id")
		}

		for _, todo := range todos {
			if todo.Id == id {
				todo.Done = true
				todoItem = &todo
				break
			}
		}

		return c.JSON(todoItem)
	})

	app.Post("/api/todo", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		todo.Id = len(todos) + 1
		todos = append(todos, *todo)

		return c.JSON(todo)
	})

	log.Fatal(app.Listen(":4000"))
}
