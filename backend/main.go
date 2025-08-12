// API WITHOUT DB

// WE DO THIS USING MEMORY , HENCE THE DATA GETS DELETED EVERY TIME WE RESTART

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
	app.Get("/api/todos", func(c *fiber.Ctx) error {
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

	app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id { // since ID is int and id is string , fmt.Sprint(ID) -> convrts ID to String
				todos = append(todos[:i], todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"succexx": true})
				// 1 2 3 4 5 [ i =4 ]
				// todos[:i] -> from index o to i-1; -->  1 2
				//todos[i+1:] -> from index i+1 to end --> 4 5
				// append these -> 1 2 4 5 [ hence 3 is deleted]			}

			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "Todo not found"})

	})

	log.Fatal(app.Listen(":4000"))
}
