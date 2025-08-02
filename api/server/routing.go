package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	OpenAPIDefinitionHandler(router)
	router.HandleFunc("/api/status", ApiStatus).Methods(http.MethodGet)

	return router

}

func ApiStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}
