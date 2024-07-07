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

func updatePlayer(id int, p *Player) (Player, error) {
	var up Player
	row := db.QueryRow(
		"UPDATE public.player SET player_name=$1, player_lv=$2 WHERE player_id=$3 RETURNING player_id, player_name, player_lv ;",
		p.PlayerName,
		p.PlayerLv,
		id,
	)

	err := row.Scan(&up.ID, &up.PlayerName, &up.PlayerLv)

	if err != nil {
		return Player{}, err
	}

	return up, err
}

func deletePlayer(id int) error {
	_, err := db.Exec("DELETE FROM public.player WHERE player_id=$1;", id)
	return err
}

func getMorePlayer() ([]Player, error) {
	rows, err := db.Query("SELECT player_id, player_name, player_lv FROM player;")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var players []Player
	for rows.Next() {
		var p Player
		err := rows.Scan(&p.ID, &p.PlayerName, &p.PlayerLv)

		if err != nil {
			return nil, err
		}

		players = append(players, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return players, nil
}
