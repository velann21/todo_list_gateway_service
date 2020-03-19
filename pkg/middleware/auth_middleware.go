package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/todo_list_gateway_service/pkg/Orchestraters"
	"github.com/todo_list_gateway_service/pkg/entities/responses"
	"github.com/todo_list_gateway_service/pkg/helpers"
	"github.com/todo_list_gateway_service/pkg/permissions"
	"net/http"
	"strconv"
)

func Authentication() mux.MiddlewareFunc {
	return func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request)  {
			logrus.WithField("EventType", "Authentication").WithField("Action","Request").Info("Authentication Start")
			w.Header().Set("Content-Type", "application/json")
			route := req.URL.Path
            resp, err := validate(route, req)
			if err != nil{
				logrus.WithField("EventType", "Authentication").WithError(err).Error("validate Failed")
				logrus.WithField("EventType", "Authentication").WithField("Status", resp.Status).WithError(err).Error("validate Failed")
				res, err := strconv.Atoi(resp.Status)
				if err != nil{
					return
				}
				w.WriteHeader(res)
				data, err := json.Marshal(resp)
				if err != nil {
					logrus.WithField("EventType", "Authentication").WithError(err).Error("Marshal Failed")
					fmt.Println(err.Error())
					return
				}
				_, err = w.Write(data)
				return
			}
			sCode, err := strconv.Atoi(resp.Status)
            if err != nil{
				logrus.WithField("EventType", "Authentication").WithError(err).Error("Atoi Failed")
				w.WriteHeader(sCode)
				data, err := json.Marshal(resp)
				if err != nil {
					logrus.WithField("EventType", "Authentication").WithError(err).Error("Marshal Failed")
					fmt.Println(err.Error())
					return
				}
				_, err = w.Write(data)
				return
			}

			handler.ServeHTTP(w, req)
			logrus.WithField("EventType", "Authentication").WithField("Action","Request").Info("Authentication Ends")
			return
		})}}


func validate(route string, req *http.Request)(responses.Response,  error){
	permission := permissions.GetPermission()
	isRoutePresent := permission[route]
	dataArr := make([]map[string]interface{}, 0)
	metaArr := make(map[string]interface{}, 0)
	if isRoutePresent != "" {
		if isRoutePresent == permissions.PUBLIC{
			return responses.Response{"200", dataArr ,metaArr}, nil
		}
		resp, err := Orchestraters.AuthServiceValidateTokenOrchestrator(req)
		if err != nil{
			return responses.Response{"500", dataArr ,metaArr},  helpers.SomethingWrong
		}
		if resp.StatusCode == 401{
			return responses.Response{"401", dataArr ,metaArr},  helpers.UnAuthorized
		}
		if resp.StatusCode == 200{
			return responses.Response{"200", dataArr ,metaArr},  nil
		}
	}else {
		return responses.Response{"404", dataArr ,metaArr},  helpers.InvalidRequest
	}

	return responses.Response{"200", dataArr ,metaArr},nil

}