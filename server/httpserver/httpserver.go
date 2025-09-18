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
	// Add more handlers here as needed
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
