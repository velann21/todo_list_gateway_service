package Orchestraters
import (
	"context"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
)
type GrpcActivityManagerService struct {
	ActivityManager proto.TodoActivityManagerClient
}

func (am GrpcActivityManagerService) CreateTodoList(ctx context.Context, request *proto.CreateTodoListRequest)(*proto.CreateTodoListResponse,error){
	resp, err := am.ActivityManager.CreateTodo(ctx, request)
	if err!=nil{
		return nil, err
	}
	return resp, nil
}