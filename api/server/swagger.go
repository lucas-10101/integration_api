package server

import (
	"api/resources"
	"net/http"

	_ "embed"

	"github.com/gorilla/mux"
)

func OpenAPIDefinitionHandler(router *mux.Router) {

	pathPrefix := router.PathPrefix("/swagger")

	pathPrefix.Handler(http.FileServerFS(resources.SwaggerFS)).Methods(http.MethodGet)
}
