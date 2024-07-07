package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

var db *sql.DB

type Player struct {
	ID         int
	PlayerName string
	PlayerLv   int
}

func main() {
	// Connection string
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName)

	sdb, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	db = sdb

	defer db.Close()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	//Connect Database
	fmt.Println("Successfully connected!")

	//Create Player
	// err = createPlayer(&Player{PlayerName: "Tanjiro", PlayerLv: 85})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Player created successfully!")

	//Get Player
	player, err := getPlayer(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Here ur Player : ", player)

	//Update Player
	// updatePlayer, err := updatePlayer(4, &Player{PlayerName: "Maprang", PlayerLv: 88})
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Updated Player : ", updatePlayer)

	//Delete Player
	// err = deletePlayer(4)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Player Delete Successfully!")

	//Get More Player
	players, err := getMorePlayer()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Here All Players : ", players)

}
