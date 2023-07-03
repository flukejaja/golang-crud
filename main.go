package main

import (
	"encoding/json"
	"io"
	"os"
	"github.com/gofiber/fiber/v2"
)

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
	} `json:"social"`
}


func main() {
    app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
        jsonFile, err := os.Open("user.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		var data map[string][]User
		err = decoder.Decode(&data)
		if err != nil && err != io.EOF {
			return err
		}
		return c.JSON(data)
    })
	app.Post("/add", func(c *fiber.Ctx) error {
        jsonFile, err := os.Open("user.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		var data map[string][]User

		err = decoder.Decode(&data)
		if err != nil && err != io.EOF {
			return err
		}
		req := new(User)
		if err := c.BodyParser(req); err != nil {
			c.Status(400).JSON(&fiber.Map{
				"status":400,
			})
		}

		data["users"] = append(data["users"], *req)

		return c.JSON(data)
    })
	app.Put("/update", func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("user.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		var data map[string][]User
		err = decoder.Decode(&data)
		req := new(User)
		if err := c.BodyParser(req); err != nil {
			c.Status(400).JSON(&fiber.Map{
				"status":400,
			})
		}

		for i, v := range data["users"] {
			if v.Name == req.Name {
				data["users"][i] = *req
			}
		}

		return c.JSON(data)
	})
	app.Delete("/delete", func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("user.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		decoder := json.NewDecoder(jsonFile)
		var data map[string][]User
		err = decoder.Decode(&data)
		req := new(User)
		if err := c.BodyParser(req); err != nil {
			c.Status(400).JSON(&fiber.Map{
				"status":400,
			})
		}
		
		for i , v := range data["users"] {
			if v.Name == req.Name {
			data["users"] = append(data["users"][:i], data["users"][i+1:]...)
			}
		}
		return c.JSON(data)
	})
    app.Listen(":3000")
}