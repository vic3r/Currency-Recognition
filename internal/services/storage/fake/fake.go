package fake

import (
	"github.com/vic3r/Currency-Recognition/internal/services/storage"
)

const ID = "fake"

type Fake struct {
	config storage.Config
}

var _ storage.Service = &Fake{}

// New creates a new mongo session
func New(config storage.Config) (storage.Service, error) {

	return nil, nil
}

// CreateExpense create a fake expense
func (f *Fake) CreateExpense(body []byte) ([]byte, error) {

	return nil, nil
}

// GetExpenses return fake expenses
func (f *Fake) GetExpenses() (string, error) {

	return "", nil
}
