package Orchestraters
import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	proto "github.com/velann21/todo_list_activity_manager/pkg/proto"
)
type GrpcActivityManagerService struct {
	ActivityManager proto.TodoActivityManagerClient
}

func (am GrpcActivityManagerService) CreateTodoList(ctx context.Context, request *proto.CreateTodoListRequest)(*proto.CreateTodoListResponse,error){
	fmt.Println(am.ActivityManager)
	resp, err := am.ActivityManager.CreateTodo(ctx, request)
	if err!=nil{
		logrus.Error("Somthing wrong:  ", err)
		return nil, err
	}
	return resp, nil
}