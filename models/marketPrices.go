package models

import "time"

type MarkerPrices struct {
	Id            uint      `json:"-"`
	AdjustedPrice float64   `json:"adjusted_price"`
	AveragePrint  float64   `json:"average_price"`
	TypeId        uint      `json:"type_id"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}
