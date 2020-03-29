package routes

import (
	"github.com/gorilla/mux"
	"github.com/todo_list_gateway_service/pkg/controller"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
)


type Configuration struct {
	ActivityManager proto.TodoActivityManagerClient
}
func (conf *Configuration)Routes(routes *mux.Router){
	routes.PathPrefix("/api/v1/todo/create").HandlerFunc(NewController(conf).CreateTodoController)
	routes.PathPrefix("/api/v1/users/").HandlerFunc(controller.UserServicesController)
	routes.PathPrefix("/api/v1/todo/").HandlerFunc(controller.TodoServicesController)
}

func NewController(configuration *Configuration) controller.GrpcController{
	grpcController := controller.GrpcController{configuration.ActivityManager}
	return grpcController
}




