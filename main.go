package main

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Player struct {
	ID         int    `json:"id"`
	PlayerName string `json:"name"`
	PlayerLv   int    `json:"level"`
}

func main() {
	// Connection string
	client, err := ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer client.Db.Close()

	//FiberGO
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Listen(":8080")
}

func getPlayerHandler(c *fiber.Ctx) error {
	// playerId := c.Params("id")
	// player, err := GetPlayer(playerId)
	return nil
}

//FuncZone
