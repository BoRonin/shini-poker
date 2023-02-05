package models

import (
	"context"
    "fmt"
    "shini/storage/postgres"
	"time"
)

type Game struct {
	Id         uint      `json:"id"`
	Title      string    `json:"title"`
    CreatedAt  time.Time `json:"created_at,omitempty"`
	Multiplier int       `json:"multiplier"`
}

func (g *Game) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	var id uint
	stmt := `insert into game (title, created_at, multiplier) values ($1, $2, $3) returning id`
	if err := postgres.DB.QueryRowContext(ctx, stmt, g.Title, g.CreatedAt, g.Multiplier).Scan(&id); err != nil {
		return err
	}
	g.Id = id
	return nil
}


func (g *Game) GetPlayers() ([]Player, error) {
    fmt.Println("Inside getplayers function")
    ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
    defer cancel()
    var players []Player

    statement := `select users.id, users.name, players.chips from users
join players on users.id = players.user_id
where players.game_id = $1`
    rows, err := postgres.DB.QueryContext(ctx, statement, g.Id)
    if err != nil {
        return nil, err
    }

    for rows.Next(){
        println("In rows")
        var player Player
        if err := rows.Scan(&player.User.Id, &player.User.Name, &player.Chips); err != nil {
            println("error in rows scan")
            println(err.Error())
            return nil, err
        }

        players = append(players, player)
     }
     return players, nil
}

func (g *Game) Finish() error {
	// Apply finished state to DB and count money
	return nil
}
