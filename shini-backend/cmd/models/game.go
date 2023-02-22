package models

import (
	"context"
	"log"
	"shini/storage/postgres"
	"time"
)

type Game struct {
	Id         uint          `json:"id"`
	Title      string        `json:"title"`
	CreatedAt  time.Time     `json:"created_at,omitempty"`
	Multiplier int           `json:"multiplier"`
	Finished   bool          `json:"is_finished"`
	Duration   time.Duration `json:"duration"`
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
where players.game_id = $1
order by players.id`
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

func (g *Game) IsFinished() error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := `select is_finished from game
	where id = $1`
	if err := postgres.DB.QueryRow(ctx, q, g.Id).Scan(&g.Finished); err != nil {
		return err
	}
	return nil
}

func GetGames() ([]Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := "select * from game order by created_at"
	rows, err := postgres.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	var games []Game
	for rows.Next() {
		var game Game
		if err := rows.Scan(&game.Id, &game.Title, &game.CreatedAt, &game.Multiplier, &game.Finished); err != nil {
			return nil, err
		}
		games = append(games, game)

	}
	q = `select age(wins.created_at, game.created_at) as duration from wins
	join players on players.id = wins.player
	join game on game.id = players.game_id
	where game.id = $1
	group by game.id, duration
	order by duration desc
	limit 1`
	for i, v := range games {
		if v.Finished {
			if err := postgres.DB.QueryRow(ctx, q, v.Id).Scan(&games[i].Duration); err != nil {
				log.Println(err)
			}
			games[i].Duration = time.Duration(games[i].Duration.Milliseconds())

		}
	}
	return games, nil
}
