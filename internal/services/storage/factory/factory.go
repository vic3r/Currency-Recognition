package factory

import (
	"errors"
	"fmt"

	"github.com/vic3r/Currency-Recognition/internal/services/storage"
	"github.com/vic3r/Currency-Recognition/internal/services/storage/fake"
	"github.com/vic3r/Currency-Recognition/internal/services/storage/mongo"
)

type factory func(storage.Config) (storage.Service, error)

var implementations = map[string]factory{
	mongo.ID: mongo.New,
	fake.ID:  fake.New,
}

// Create returns a new payment service with a given configuration
func Create(rawConfig map[string]interface{}) (storage.Service, error) {
	implementation, ok := rawConfig["implementation"].(string)
	if !ok {
		return nil, errors.New("invalid configuration: implementation is not a string")
	}

	serviceFactory, ok := implementations[implementation]
	if !ok {
		return nil, fmt.Errorf("job service implementation not found: %v", implementation)
	}

	return serviceFactory(rawConfig)
}
