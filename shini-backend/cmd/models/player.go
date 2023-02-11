package models

import (
	"context"
	"fmt"
	"shini/storage/postgres"
)

type Player struct {
	Id         uint `json:"id,omitempty"`
	User       User `json:"user"`
	Game       Game `json:"game,omitempty"`
	Chips      int  `json:"chips,omitempty"`
	ChipsFinal int  `json:"chips_final,omitempty"`
}
type PlayersForList struct {
	Id    uint   `json:"id"`
	Name  string `json:"user_name"`
	Chips int    `json:"chips"`
	Login string `json:"login"`
}

func (p *Player) AddChips(number int) error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	stmt := `update players 
	set chips = (chips + $1)
	where id = $2 returning chips`
	if err := postgres.DB.QueryRow(ctx, stmt, number, p.Id).Scan(&p.Chips); err != nil {
		return err
	}
	return nil
}

func (p *Player) Create() error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	stmt := "insert into players (user_id, game_id, chips) values ($1, $2, $3) returning id"
	var id uint
	err := postgres.DB.QueryRow(ctx, stmt, p.User.Id, p.Game.Id, p.Chips).Scan(&id)
	if err != nil {
		return err
	}
	p.Id = id
	return nil
}

func (p *Player) Win(winId int) error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := "select is_finished from game where id = $1"
	var finished bool
	err := postgres.DB.QueryRow(ctx, q, p.Game.Id).Scan(&finished)
	if err != nil {
		return err
	}
	if finished {
		return fmt.Errorf("the game %d is finished", p.Game.Id)
	}
	q = "insert into wins (player, combination) values ($1, $2)"
	_, err = postgres.DB.Exec(ctx, q, p.Id, winId)
	if err != nil {
		return err
	}
	return nil
}
func (p *Player) SetFinalChips() (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := `update players set chips_final = $1 where id = $2 
	returning (chips_final - chips) * (select game.multiplier from game                                  
									  where id = $3)`
	var score int
	err := postgres.DB.QueryRow(ctx, q, p.ChipsFinal, p.Id, p.Game.Id).Scan(&score)
	if err != nil {
		return 0, err
	}
	return score, nil
}
