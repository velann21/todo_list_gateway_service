package requests

import (
	"encoding/json"
	"github.com/todo_list_gateway_service/pkg/helpers"
	protoTodo "github.com/velann21/todo_list_activity_manager/pkg/proto"
	"io"
)

//type Notification struct {
//	Email   Email   `json:"email"`
//	Message Message `json:"message"`
//}
//
//type Email struct {
//	Notification bool   `json:"notification"`
//	EmailID      string `json:"emailID"`
//}
//
//type Message struct {
//	Notification bool   `json:"notification"`
//	CountryCode  string `json:"countryCode"`
//	MobileNumber string `json:"mobileNumber"`
//}
//
//type TodoRequest struct {
//	UserID        string       `json:"userID"`
//	Name          string       `json:"name"`
//	Description   string       `json:"description"`
//	Notification  Notification `json:"notification"`
//	DueDate       time.Time    `json:"dueDate"`
//	CustomDueDate string       `json:"customDueDate"`
//	Repeatmode    bool         `json:"repeatMode"`
//	Tag           string       `json:"tag"`
//	SubTasks      []SubTask    `json:"subtask"`
//}
//
//type SubTask struct {
//	Name        string `json:"name"`
//	Status      bool   `json:"status"`
//	Description string `json:"description"`
//	Offset      int    `json:"offset"`
//}
//
//func (todoRequest *TodoRequest) PopulateTodoRequest(body io.ReadCloser) error {
//	if body == nil{
//		return helpers.InvalidRequest
//	}
//	todo := &protoTodo.CreateTodoListRequest{}
//	decoder := json.NewDecoder(body)
//	err := decoder.Decode(todo)
//	if err != nil {
//		return helpers.InvalidRequest
//	}
//	return nil
//}


func PopulateCreateTodo(body io.Reader) (*protoTodo.CreateTodoListRequest, error) {
	req := &protoTodo.CreateTodoListRequest{}
	decode := json.NewDecoder(body)
	err := decode.Decode(req)
	if err != nil{
		return nil, helpers.InvalidRequest
	}
	return req, nil
}
