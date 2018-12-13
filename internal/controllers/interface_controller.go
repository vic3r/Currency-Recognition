package controllers

import (
	"net/http"
)

type job interface {
	GetExpenses(w http.ResponseWriter, r *http.Request)
	PostExpense(w http.ResponseWriter, r *http.Request)
}
