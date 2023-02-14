package models

import (
	"context"
	"shini/storage/postgres"
)

type Stat struct {
	UserName string `json:"user"`
	Chips    int    `json:"chips"`
}

func GetStats() ([]Stat, error) {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := `select name, sum(players.chips_final) from users
	join players on players.user_id = users.id
	join game on players.game_id = game.id
	group by name
	order by sum(players.chips_final) desc`
	rows, err := postgres.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	var stats []Stat
	for rows.Next() {
		newstat := new(Stat)
		if err := rows.Scan(&newstat.UserName, &newstat.Chips); err != nil {
			return nil, err
		}
		stats = append(stats, *newstat)
	}
	return stats, nil
}
