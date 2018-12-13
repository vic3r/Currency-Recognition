package storage

// Service is a job service interface
type Config map[string]interface{}

type Service interface {
	CreateExpense([]byte) ([]byte, error)
	GetExpenses() (string, error)
}
