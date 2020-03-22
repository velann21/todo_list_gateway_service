package routes

import (
	"github.com/gorilla/mux"
	"github.com/todo_list_gateway_service/pkg/controller"
)


func Routes(routes *mux.Router){
	routes.PathPrefix("/api/v1/users/").HandlerFunc(controller.UserServicesController)
	routes.PathPrefix("/api/v1/todo/").HandlerFunc(controller.TodoServicesController)
}




