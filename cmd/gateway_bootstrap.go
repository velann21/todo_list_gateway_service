package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_gateway_service/pkg/middleware"
	"github.com/todo_list_gateway_service/pkg/routes"
	"log"
	"net/http"
	"time"
)

func main(){
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat:time.RFC3339,})
    //helpers.SetEnv()
	r := mux.NewRouter().StrictSlash(false)
	r.Use(middleware.TraceLogger())
	r.Use(middleware.Authentication())
	routes.Routes(r)
	logrus.WithField("EventType", "Bootup").Info("Booting up server at port : "+"8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.WithField("EventType", "Server Bootup").WithError(err).Error("Server Bootup Error")
		log.Fatal(err)
	}
}

