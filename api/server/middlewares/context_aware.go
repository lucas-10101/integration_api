package middlewares

import (
	"api/services"
	"context"
	"net/http"
	"time"
)

type ContextKey string

type RequestContext struct {
	QueryParameters services.UrlQueryParser
}

type ContextLoader struct{}

var (
	QueryParameters  ContextKey = "query_parameters"
	SQLConnection    ContextKey = "sql_connection"
	NoSQLConnection  ContextKey = "nosql_connection"
	RequestStartTime ContextKey = "request_start_time"
	Language         ContextKey = "language"
)

func (m *ContextLoader) Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(ResponseWriter http.ResponseWriter, request *http.Request) {

		var newContext context.Context

		newContext = context.WithValue(request.Context(), RequestStartTime, time.Now())
		newContext = context.WithValue(newContext, QueryParameters, &services.UrlQueryParser{Values: request.URL.Query()})

		request = request.WithContext(newContext)
		handler.ServeHTTP(ResponseWriter, request)
	})
}
