package models

import (
	"context"
	"shini/storage/postgres"
)

type Combination struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Img  string `json:"image_url"`
}

func GetAllCombinations(c *[]Combination) error {
	ctx, cancel := context.WithTimeout(context.Background(), postgres.DBTimeout)
	defer cancel()
	q := `select id, name from combinations`
	rows, err := postgres.DB.Query(ctx, q)
	if err != nil {
		return err
	}
	for rows.Next() {
		var comb Combination
		if err := rows.Scan(&comb.Id, &comb.Name); err != nil {
			return err
		}
		*c = append(*c, comb)
	}
	return nil
}
