package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	router := chi.NewRouter()

	log.Println("Serving on :8080. Example here: http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", router), nil)
}
