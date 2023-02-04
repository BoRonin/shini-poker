package models

import (
    "context"
    "shini/storage/postgres"
)

type Player struct {
    Id uint `json:"id"`
    User User `json:"user"`
    Game Game `json:"game"`
    Chips int `json:"chips"`
}

func (p *Player) AddChips(number int) {

}

func (p *Player) Create() error {
    ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
    defer cancel()
    stmt := "insert into players (user_id, game, chips) values ($1, $2, $3) returning id"
    var id uint
    if err := postgres.DB.QueryRowContext(ctx, stmt, p.User.Id, p.Game.Id, p.Chips).Scan(&id); err != nil {
        return err
    }
    p.Id = id
    return nil
}