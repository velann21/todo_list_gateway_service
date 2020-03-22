package controller

import (
	"github.com/sirupsen/logrus"
	orchestrator "github.com/todo_list_gateway_service/pkg/Orchestraters"
	"github.com/todo_list_gateway_service/pkg/entities/responses"
	"net/http"
)


func UserServicesController(response http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithField("Action","Request").Info("UserServicesController Start")
	resp, err := orchestrator.UserServiceOrchestrator(req)
	finalResp := responses.Response{}
	if err != nil{
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("UserServiceOrchestrator Failed")
		if resp != nil{
			finalResp.Response(response, resp.StatusCode, resp.Body)
			return
		}
		return
	}
	finalResp.Response(response, resp.StatusCode, resp.Body)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithField("Action","Response").Info("UserServicesController Start")
	return
}


func TodoServicesController(response http.ResponseWriter, req *http.Request) {
	eventType, traceID := getEventTypeAndTraceID(req)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithField("Action","Request").Info("TodoServicesController Start")
	resp, err := orchestrator.TodoServiceOrchestrator(req)
	finalResp := responses.Response{}
	if err != nil{
		logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithError(err).Error("TodoServicesController Failed")
		if resp != nil{
			finalResp.Response(response, resp.StatusCode, resp.Body)
			return
		}
		return
	}
	finalResp.Response(response, resp.StatusCode, resp.Body)
	logrus.WithField("EventType", eventType).WithField("TraceID", traceID).WithField("Action","Response").Info("TodoServicesController Start")
	return
}

func getEventTypeAndTraceID(req *http.Request)(string, string){
	eventType := req.Header.Get("X-EventType")
	traceID := req.Header.Get("X-TraceID")
	return eventType, traceID
}