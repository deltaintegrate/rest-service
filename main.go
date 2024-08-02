package main

import (
	"net/http"
	"rest-service/handlers"
	"rest-service/utils"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	router := mux.NewRouter()
	handlers.RegisterHolidayRoutes(router)
	utils.InitLogger()

	utils.Log.Info("Starting server on :8080")
	logrus.Fatal(http.ListenAndServe(":8080", router))
}
