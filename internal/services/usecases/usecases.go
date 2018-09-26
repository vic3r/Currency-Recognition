package usecases

import (
	"context"

	"github.com/vic3r/Currency-Recognition/internal/schema"
	"github.com/vic3r/Currency-Recognition/internal/services"
)

//UseCases is a usecases imp for all the endpoints
type UseCases struct {
	conf services.Config
}

var _ services.ServiceRepository = &UseCases{}

// Close closes a connection
func (u *UseCases) Close() {

}

// InsertCoin is a usecases imp for inserting a coin
func (u *UseCases) InsertCoin(ctx context.Context, coin schema.Coin) error {
	return u.InsertCoin(ctx, coin)
}

// ListCoins is a usecases imp for listing coins
func (u *UseCases) ListCoins(ctx context.Context, skip uint64, take uint64) ([]schema.Coin, error) {
	return u.ListCoins(ctx, skip, take)
}
