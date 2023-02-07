package models

import (
	"context"
	"shini/storage/postgres"
)

type Player struct {
	Id         uint `json:"id,omitempty"`
	User       User `json:"user"`
	Game       Game `json:"game,omitempty"`
	Chips      int  `json:"chips,omitempty"`
	ChipsFinal int  `json:"chips_final,omitempty"`
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
	q := "insert into wins (player, combination) values ($1, $2)"
	_, err := postgres.DB.Exec(ctx, q, p.Id, winId)
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
