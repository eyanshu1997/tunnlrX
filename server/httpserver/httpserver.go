package httpserver

import (
	"fmt"
	"net/http"
)

type handleFuncEntry struct {
	pattern string
	handler func(http.ResponseWriter, *http.Request)
}

var handleFuncList = []handleFuncEntry{
	{pattern: "/health", handler: func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}},
	// Add a catch all handler
	{pattern: "/", handler: proxyHandler},
}

type HttpServer struct {
	httpMux *http.ServeMux
	port    int
}

func NewHttpServer(port int) *HttpServer {
	httpMux := http.NewServeMux()
	for _, handleFuncEntry := range handleFuncList {
		httpMux.HandleFunc(handleFuncEntry.pattern, handleFuncEntry.handler)
	}
	return &HttpServer{
		httpMux: httpMux,
		port:    int(port),
	}
}

func (s *HttpServer) Start() error {
	return http.ListenAndServe(fmt.Sprintf(":%d", s.port), s.httpMux)
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// lets read the host name of the request
	host := r.Host
	fmt.Fprintf(w, "Hello from proxy handler, you requested for host: %s", host)
	w.WriteHeader(http.StatusOK)
}
