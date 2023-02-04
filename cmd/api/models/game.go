package models

import (
    "context"
    "shini/storage/postgres"
    "time"
)

type Game struct {
    Id uint `json:"id"`
    Title string `json:"title"`
    CreatedAt  time.Time `json:"created_at"`
    Multiplier int    `json:"multiplier"`
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

func (g *Game) Finish() error {
    // Apply finished state to DB and count money
    return nil
}



