package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vic3r/Currency-Recognition/internal/services/storage"
)

// JobController is a job defined job struct
type JobController struct {
	Storage storage.Service
}

var _ job = &JobController{}

// GetExpenses retrieves all the expenses
func (c *JobController) GetExpenses(w http.ResponseWriter, r *http.Request) {
	response, err := c.Storage.GetExpenses()
	if err != nil {
		fmt.Printf("Failed to retrieve expenses")
	}

	fmt.Fprintf(w, response)
}

// PostExpense creates an expense
func (c *JobController) PostExpense(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Printf("Not possible to read request")
	}

	response, err := c.Storage.CreateExpense(body)
	if err != nil {
		fmt.Printf("Failed to create expense")
	}

	w.Header().Set("content-type", "application/json")
	w.Write(response)
}
