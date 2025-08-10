package main

import (
	"fmt" // fmt stands for fromatted iinput /output
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() { // main funtion -> first executed

	fmt.Print("Hello worlds 123 ")

	app := fiber.New() // starts fibre the web framework of GO // go run main.go --> in terminal to run it
    

	app.Get("/",func(c * fiber.Ctx)error{
		return c.Status(200).JSON(fiber.Map{"msg":"hello world"})
	})

	log.Fatal(app.Listen(":4000")) // listen to port 4000 // log.Fatal is like check for errors if existing 

}
