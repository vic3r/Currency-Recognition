package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"github.com/vic3r/Currency-Recognition/internal/controllers"
	storageFactory "github.com/vic3r/Currency-Recognition/internal/services/storage/factory"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("$HOME/go/src/github.com/vic3r/Currency-Recognition/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed reading config: %v", err)
	}

	storageConfig := viper.Get("services.storage").(map[string]interface{})
	storage, err := storageFactory.Create(storageConfig)
	if err != nil {
		log.Fatalln("error creating vehicle service:", err)
	}

	controller := &controllers.JobController{Storage: storage}
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/GetExpenses", controller.GetExpenses).Methods("GET")
	router.HandleFunc("/CreateExpense", controller.PostExpense).Methods("POST")

	apiConfig := viper.Get("api").(map[string]interface{})
	host := apiConfig["host"].(string)
	port := apiConfig["port"].(string)
	log.Fatal(http.ListenAndServe(host+":"+port, router))

}
