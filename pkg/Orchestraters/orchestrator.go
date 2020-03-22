package Orchestraters

import (
	"github.com/sirupsen/logrus"
	"github.com/todo_list_gateway_service/pkg/helpers"
	"io/ioutil"
	"net/http"
)

func UserServiceOrchestrator(req *http.Request) (*http.Response, error) {
	url := req.URL
	body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		return nil, err
	}
	logrus.Info("Inside the ")
	logrus.Info(helpers.ReadEnv(helpers.USERSRV_HOST)+url.Path+"?"+url.Query().Encode())
	resp , err := helpers.HttpRequest(req.Method, req, helpers.ReadEnv(helpers.USERSRV_HOST)+url.Path+"?"+url.Query().Encode(), body)
	if err!= nil{
       return nil, err
	}
	return  resp, nil
}

func TodoServiceOrchestrator(req *http.Request) (*http.Response, error) {
	url := req.URL
	body, err := ioutil.ReadAll(req.Body)
	if err != nil{
		return nil, err
	}
	resp , err := helpers.HttpRequest(req.Method, req, helpers.ReadEnv(helpers.TODOSRV_HOST)+url.Path+"?"+url.Query().Encode(), body)
	if err!= nil{
		return nil, err
	}
	return  resp, nil
}


func AuthServiceValidateTokenOrchestrator(req *http.Request) (*http.Response, error) {
	resp , err := helpers.HttpRequest("GET", req, helpers.ReadEnv(helpers.AUTHSRV_HOST)+"/api/v1/auth/authenticate", nil)
	if err!= nil{
		return nil, err
	}
	return  resp, nil
}











//type Host struct {
//	hostname string
//	scheme   string
//	httpsEnabled bool
//}
//
//func WithHost(hostname string) *Host {
//	return &Host{hostname: hostname, scheme: "http"}
//}
//
//func (h *Host) IsHttpsEnabled(enabled bool) *Host {
//	h.httpsEnabled = enabled
//	return h
//}
//
//func (h *Host) SetScheme(scheme string) *Host {
//	h.scheme = scheme
//	return h
//}
//
//type roundTripper func(*http.Request) (*http.Response, error)
//
//func (f roundTripper) RoundTrip(req *http.Request) (*http.Response, error) { return f(req) }
//
//type OldHost struct {
//	oldHost string
//}
//
//func (h *OldHost) rt(req *http.Request) (*http.Response, error) {
//	req.Header.Set("Host", h.oldHost) // <--- I set it here as well
//	return http.DefaultTransport.RoundTrip(req)
//}
//

//h = WithHost("localhost:7000")
//var oldHost = req.Host
//
//appURL := h.scheme + "://" + h.hostname
//
//u, err := url.Parse(appURL)
//if err != nil {
//
//}
//
//var (
//	proto string
//	forwardingPort string
//)
//
//if h.httpsEnabled {
//	proto = "https"
//	forwardingPort = "443"
//} else {
//	proto = "http"
//	forwardingPort = "80"
//}
//
//director := func(req *http.Request) {
//
//	req.Header.Set("X-GoProxy", "RCControllerProxy")
//	req.Header.Set("X-Forwarded-Host", oldHost)
//	req.Header.Set("X-Real-IP", req.RemoteAddr)
//	req.Header.Set("X-Forwarded-Proto", proto )
//	req.Header.Set("X-Forwarded-For", req.RemoteAddr)
//	req.Header.Set("Host", oldHost)
//	req.Header.Set("X-Forwarded-Port", forwardingPort)
//
//	req.URL.Scheme = u.Scheme
//	req.URL.Host = u.Host
//	req.Host = ""
//}
//
//host := &OldHost{oldHost: oldHost}
//
//proxy := &httputil.ReverseProxy{
//	Director:  director,
//	Transport: roundTripper(host.rt),
//}
//response.WriteHeader(400)
//proxy.ServeHTTP(response, req)


