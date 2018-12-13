package mongo

import (
	"encoding/json"

	"github.com/vic3r/Currency-Recognition/internal/services/models"
	"github.com/vic3r/Currency-Recognition/internal/services/storage"
	"gopkg.in/mgo.v2/bson"
)

const ID = "mongo"

// Mongo is a mongo struct config
type Mongo struct {
	config storage.Config
}

var _ storage.Service = &Mongo{}

// New creates a new mongo session
func New(config storage.Config) (storage.Service, error) {
	err := connect(config)
	if err != nil {
		return nil, err
	}

	return &Mongo{config}, nil
}

// CreateExpense create a expense
func (m *Mongo) CreateExpense(body []byte) ([]byte, error) {
	col := db.DB(database).C(collection)
	var expense models.Expense
	err := json.Unmarshal(body, &expense)
	if err != nil {
		return []byte(""), err
	}

	err = col.Insert(expense)
	if err != nil {
		return []byte(""), err
	}

	expenseResponse, err := json.Marshal(expense)
	if err != nil {
		return []byte(""), err
	}

	return expenseResponse, nil
}

// GetExpenses return all expenses
func (m *Mongo) GetExpenses() (string, error) {
	col := db.DB(database).C(collection)
	results := []models.Expense{}
	col.Find(bson.M{"title": bson.RegEx{"", ""}}).All(&results)

	jsonString, err := json.Marshal(results)
	if err != nil {
		return "", err
	}

	return string(jsonString), nil
}
