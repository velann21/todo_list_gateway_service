package helpers

import (
	"bytes"
	"net/http"
	"os"
	"time"
)

const (
	USERSRV_HOST = "USERSRV"
    AUTHSRV_HOST = "AUTHSRV"
)
func HttpRequest(methodType string,actualReq *http.Request, Url string,body []byte) (*http.Response, error){
	req, err := http.NewRequest(methodType, Url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	var client http.Client
	req.Header = actualReq.Header
	client.Timeout = 15 * time.Second
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func ReadEnv(envType string) string{
	switch envType{
	case USERSRV_HOST:
		return os.Getenv("USERSRV_HOST")
	case AUTHSRV_HOST:
		return os.Getenv("AUTHSRV_HOST")
	default:
		return ""
	}
}

func SetEnv(){
	os.Setenv("USERSRV_HOST", "http://localhost:8081")
	os.Setenv("AUTHSRV_HOST","http://localhost:8083")
}
