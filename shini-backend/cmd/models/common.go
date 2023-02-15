package models

import (
	"context"
	"shini/storage/postgres"
)

type Stat struct {
	Chips_after int `json:"chips_after"`
	Chips_taken int `json:"chips_taken"`
}
type StatNameCount struct {
	Combination string `json:"combination"`
	Count       int    `json:"count"`
}
type AllStats struct {
	Username string `json:"username"`
	Login    string `json:"login"`
	Stat
	StatsNameCount []StatNameCount `json:"name_count_stats"`
}

func GetStats() ([]AllStats, error) {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := `select name, login, sum(players.chips_final) as chips_left, sum(players.chips) as chips_taken from users
	join players on players.user_id = users.id
	join game on players.game_id = game.id
	group by name, login
	order by chips_taken desc`
	rows, err := postgres.DB.Query(ctx, q)
	if err != nil {
		return nil, err
	}
	var allStats []AllStats
	for rows.Next() {
		newstat := new(Stat)
		var name string
		var login string
		if err := rows.Scan(&name, &login, &newstat.Chips_after, &newstat.Chips_taken); err != nil {
			return nil, err
		}
		allStats = append(allStats, AllStats{Username: name, Login: login, Stat: *newstat})
	}
	q2 := `select users.name, combinations.name, coalesce(count(wins.player), 0)as w
	from combinations
	join wins on wins.combination = combinations.id
	join players on wins.player = players.id
	join users on players.user_id = users.id
	group by users.name, combinations.name
	order by users.name, w desc`
	rows2, err := postgres.DB.Query(ctx, q2)
	if err != nil {
		return nil, err
	}
	for rows2.Next() {
		newStatNC := new(StatNameCount)
		var name string
		if err := rows2.Scan(&name, &newStatNC.Combination, &newStatNC.Count); err != nil {
			return nil, err
		}
		for i, v := range allStats {
			if v.Username == name {
				allStats[i].StatsNameCount = append(allStats[i].StatsNameCount, *newStatNC)
			}
		}

	}
	return allStats, nil
}
