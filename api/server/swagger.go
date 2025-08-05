package server

import (
	"api/static"
	"net/http"

	_ "embed"

	"github.com/gorilla/mux"
)

func OpenAPIDefinitionHandler(router *mux.Router) {

	pathPrefix := router.PathPrefix("/swagger")

	pathPrefix.Handler(http.FileServerFS(static.SwaggerFS)).Methods(http.MethodGet)
}
