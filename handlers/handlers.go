package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/nelsomar/twittor/middlew"
	"github.com/nelsomar/twittor/routers"
)

/* HandlersAPI, Set port an listening server */
func HandlersAPI() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "7200"
	}

	handler := cors.AllowAll().Handler(router)

	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
