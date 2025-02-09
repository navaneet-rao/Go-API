package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"

	"github.com/navaneet-rao/Go-API/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Printf("Staring GO APT service....")

	err := http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Error(err)
	}
}
