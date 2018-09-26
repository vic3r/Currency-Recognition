package services

import (
	"context"

	"github.com/vic3r/Currency-Recognition/internal/schema"
)

// Config is a configuration for the service imp
type Config map[string]interface{}

// ServiceRepository is an interface which will be implemented by the service
type ServiceRepository interface {
	Close()
	InsertCoin(ctx context.Context, coin schema.Coin) error
	ListCoins(ctx context.Context, skip uint64, take uint64) ([]schema.Coin, error)
}
