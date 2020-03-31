package main

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_gateway_service/pkg/middleware"
	"github.com/todo_list_gateway_service/pkg/routes"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
	"log"
	"net/http"
	"os"
	"time"
)

func main(){
	logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat:time.RFC3339,})
    //helpers.SetEnv()
	r := mux.NewRouter().StrictSlash(false)
	r.Use(middleware.TraceLogger())
	r.Use(middleware.Authentication())

	resolver.SetDefaultScheme("dns")
	amConn, err := grpc.Dial(
		"todolistsrv:50051",
		grpc.WithInsecure(),
		grpc.WithBalancerName(roundrobin.Name),
	)
	//amConn, err := grpc.Dial("todolistsrv:50051", grpc.WithInsecure())
	if err != nil{
		logrus.Error("Something went wrong while calling AM grpc server")
		os.Exit(1)
	}
	amClient := proto.NewTodoActivityManagerClient(amConn)
	configuration := routes.Configuration{amClient}
	configuration.Routes(r)

	logrus.WithField("EventType", "Bootup").Info("Booting up server at port : "+"8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logrus.WithField("EventType", "Server Bootup").WithError(err).Error("Server Bootup Error")
		log.Fatal(err)
	}
}

