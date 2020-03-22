package responses

import (
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