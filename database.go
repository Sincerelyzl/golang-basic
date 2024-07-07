package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "mydatabase"
	username     = "myuser"
	password     = "mypassword"
)

type DatabaseController struct {
	Db *sql.DB
}

func ConnectDatabase() (*DatabaseController, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		username,
		password,
		databaseName,
	)

	sdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	client := &DatabaseController{
		Db: sdb,
	}

	err = client.Db.Ping()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (client *DatabaseController) CreatePlayer(p *Player) error {
	_, err := client.Db.Exec(
		"INSERT INTO public.player(player_name, player_lv)	VALUES ($1, $2)",
		p.PlayerName,
		p.PlayerLv,
	)

	return err
}

func (client *DatabaseController) GetPlayer(id int) (Player, error) {
	var p Player
	row := client.Db.QueryRow("SELECT player_id, player_name, player_lv FROM player WHERE player_id=$1;", id)

	err := row.Scan(&p.ID, &p.PlayerName, &p.PlayerLv)

	if err != nil {
		return Player{}, err
	}
	return p, nil
}

func (client *DatabaseController) UpdatePlayer(id int, p *Player) (Player, error) {
	var up Player
	row := client.Db.QueryRow(
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

func (client *DatabaseController) DeletePlayer(id int) error {
	_, err := client.Db.Exec("DELETE FROM public.player WHERE player_id=$1;", id)
	return err
}

func (client *DatabaseController) GetMorePlayer() ([]Player, error) {
	rows, err := client.Db.Query("SELECT player_id, player_name, player_lv FROM player;")

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
