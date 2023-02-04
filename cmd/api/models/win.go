package models

import "time"

type Win struct {
    Id        uint      `json:"id"`
    CreatedAt time.Time `json:"created_at"`
    Player Player `json:"player"`
    Combination Combination `json:"combination"`
}