package responses

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Response struct
type Response struct {
	Status string                   `json:"status"`
	Data   []map[string]interface{} `json:"data"`
	Meta   map[string]interface{}   `json:"meta,omitempty"`
}

func (entity *Response) NewActivityRegisterResp(id *string) {
	responseData := make([]map[string]interface{}, 0)
	data := make(map[string]interface{})
	data["id"] = *id
	responseData = append(responseData, data)
	entity.Data = responseData
	metaData := make(map[string]interface{})
	metaData["message"] = "User task created"
}

func (entity *Response) Response(response http.ResponseWriter, statusCode int, reader io.Reader){
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	body, err := ioutil.ReadAll(reader)
		if err != nil {
		fmt.Println(err.Error())
		return
	}
	_, err = response.Write(body)
	return
}


// SendResponse send http response
func (entity *Response) SendResponse(rw http.ResponseWriter, statusCode int) {
	rw.Header().Set("Content-Type", "application/json")

	switch statusCode {
	case http.StatusOK:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	case http.StatusCreated:
		rw.WriteHeader(http.StatusCreated)
		entity.Status = http.StatusText(http.StatusCreated)
	case http.StatusAccepted:
		rw.WriteHeader(http.StatusAccepted)
		entity.Status = http.StatusText(http.StatusAccepted)
	default:
		rw.WriteHeader(http.StatusOK)
		entity.Status = http.StatusText(http.StatusOK)
	}

	// send response
	_ = json.NewEncoder(rw).Encode(entity)

	return
}
