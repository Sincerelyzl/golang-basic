package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

type Player struct {
	ID         int    `json:"id"`
	PlayerName string `json:"name"`
	PlayerLv   int    `json:"level"`
}

var Client *DatabaseController

func main() {
	// Connection string
	var err error
	Client, err = ConnectDatabase()
	if err != nil {
		panic(err)
	}
	defer Client.Db.Close()

	//FiberGO
	app := fiber.New()

	//API ZONE
	app.Get("/player/:id", getPlayerHandler)
	app.Post("/player", createPlayerHandler)
	app.Put("/player/:id", updatePlayerHandler)
	app.Delete("/player/:id", deletePlayerHandler)
	app.Get("/player", getAllPlayerHandler)

	app.Listen(":8080")
}

func getPlayerHandler(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	player, err := Client.GetPlayer(playerId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(player)

}

func createPlayerHandler(c *fiber.Ctx) error {
	p := new(Player)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := Client.CreatePlayer(p)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(p)
}

func updatePlayerHandler(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	p := new(Player)
	if err := c.BodyParser(p); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	player, err := Client.UpdatePlayer(playerId, p)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(player)
}

func deletePlayerHandler(c *fiber.Ctx) error {
	playerId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = Client.DeletePlayer(playerId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusNoContent)

}

func getAllPlayerHandler(c *fiber.Ctx) error {
	players, err := Client.GetMorePlayer()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.JSON(players)
}
