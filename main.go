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

}

func createPlayer(p *Player) error {
	_, err := db.Exec(
		"INSERT INTO public.player(player_name, player_lv)	VALUES ($1, $2)",
		p.PlayerName,
		p.PlayerLv,
	)

	return err
}

func getPlayer(id int) (Player, error) {
	var p Player
	row := db.QueryRow("SELECT player_id, player_name, player_lv FROM player WHERE player_id=$1;", id)

	err := row.Scan(&p.ID, &p.PlayerName, &p.PlayerLv)

	if err != nil {
		return Player{}, err
	}
	return p, nil
}
