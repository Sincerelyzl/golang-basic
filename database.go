package main

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
