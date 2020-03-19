package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	permission "github.com/todo_list_gateway_service/pkg/permissions"
)

// TraceLogger mux middleware function for adding headers for tracing
// Authentication() is to route the request to Authentication service to make
// Sure the request is from Authenticated person
func TraceLogger() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			logger := logrus.WithField("route", "")
			logger.Info("requested")
			logrus.Info("req.url.path ", req.URL.Path)
			eventType := permission.EventsPermission[req.URL.Path]
			if eventType == "" {
				handler.ServeHTTP(w, req)
				return
			}
			syntheticTraceID := uuid.New().String()
			req.Header.Add("X-TraceID", syntheticTraceID)
			req.Header.Add("X-EventType", eventType)
			handler.ServeHTTP(w, req)
			return
		})
	}
}
