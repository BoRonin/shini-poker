package models

import (
	"context"
	"shini/storage/postgres"
	"time"
)

type Game struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
	CreatedAt  time.Time `json:"created_at,omitempty"`
	Multiplier int       `json:"multiplier"`
	Finished   bool      `json:"is_finished"`
}

func (g *Game) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	var id uint
	stmt := `insert into game (title, created_at, multiplier) values ($1, $2, $3) returning id`
	if err := postgres.DB.QueryRow(ctx, stmt, g.Title, g.CreatedAt, g.Multiplier).Scan(&id); err != nil {
		return err
	}
	g.Id = id
	return nil
}

func (g *Game) GetPlayers() ([]PlayersForList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	var players []PlayersForList

	statement := `select players.id, users.name, players.chips, users.login from users
join players on users.id = players.user_id
where players.game_id = $1`
	rows, err := postgres.DB.Query(ctx, statement, g.Id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var player PlayersForList
		if err := rows.Scan(&player.Id, &player.Name, &player.Chips, &player.Login); err != nil {
			return nil, err
		}

		players = append(players, player)
	}
	return players, nil
}

func (g *Game) Finilize() error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := "update game set is_finished = true where id = $1"
	_, err := postgres.DB.Exec(ctx, q, g.Id)
	if err != nil {
		return err
	}
	return nil
}
