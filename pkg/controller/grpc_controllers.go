package controller

import (
	"context"
	"github.com/todo_list_gateway_service/pkg/Orchestraters"
	"github.com/todo_list_gateway_service/pkg/entities/requests"
	"github.com/todo_list_gateway_service/pkg/entities/responses"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
	"net/http"
)
type GrpcController struct {
	ActivityManager proto.TodoActivityManagerClient
}

func (grpcCtrl GrpcController) CreateTodoController(res http.ResponseWriter, req *http.Request){
	succesResp := responses.Response{}
	reqBody, err := requests.PopulateCreateTodo(req.Body)
	if err != nil{
		responses.HandleError(res, err)
		return
	}

	service := Orchestraters.GrpcActivityManagerService{grpcCtrl.ActivityManager}
	resp, err := service.CreateTodoList(context.Background(), reqBody)
	if err != nil{
		responses.HandleError(res, err)
		return
	}
	succesResp.NewActivityRegisterResp(&resp.Data[0].ID)
    succesResp.SendResponse(res, http.StatusOK)
	return
}