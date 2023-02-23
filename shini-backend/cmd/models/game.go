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

type PlayersGameStat struct{
    Id uint `json:"id"`
    Login string `json:"login"`
    Money int `json:"money"`
    Wins int `json:"wins"`

}
type GameStatsById struct {
    PlayerStats []PlayersGameStat `json:"player_stats"`
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
func (g *Game)GetStats() (GameStatsById, error) {
    ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
    defer cancel()
    q := `select users.id, users.login, coalesce((players.chips_final-players.chips)* game.multiplier, 0) as money, count(wins.id) as wins from users
join players on players.user_id = users.id
join wins on wins.player = players.id
join game on players.game_id = game.id
where game.id = $1
group by users.id, users.login, money
order by users.id
`
    rows, err := postgres.DB.Query(ctx, q, g.Id)
    if err != nil {
        return GameStatsById{}, err
    }
    var gameStat GameStatsById
    for rows.Next(){
        var playerStat PlayersGameStat
        err := rows.Scan(&playerStat.Id, &playerStat.Login, &playerStat.Money, &playerStat.Wins)
        if err != nil {
            return GameStatsById{}, err
        }
        gameStat.PlayerStats = append(gameStat.PlayerStats, playerStat)
    }
    return gameStat, nil
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
