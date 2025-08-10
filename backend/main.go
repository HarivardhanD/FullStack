package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct { // struct to store different data types for a todo
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello worlds 123")

	app := fiber.New() // starts Fiber web framework

	todos := []Todo{}

	// GET: Just to test server
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	// POST: Create a todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		if err := c.BodyParser(todo); err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Todo body is required"})
		}

		todo.ID = len(todos) + 1
		todos = append(todos, *todo)

		return c.Status(201).JSON(todos)
	})

	// PATCH: Update a todo by ID
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i := range todos {
			if fmt.Sprint(todos[i].ID) == id {
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})
	})

	log.Fatal(app.Listen(":4000"))
}
