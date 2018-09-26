package schema

import (
	"time"
)

// Coin is coin defined struct
type Coin struct {
	ID        string    `json:"id"`
	Value     string    `json:"value"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"created_at"`
}
