package rest

import (
	"context"
	"net/http"
	"time"
)

type HttpServer struct {
	httpServer *http.Server
}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (H *HttpServer) Run(port string, handler http.Handler) error {
	H.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return H.httpServer.ListenAndServe()
}

func (H *HttpServer) Shutdown(ctx context.Context) error {
	return H.httpServer.Shutdown(ctx)
}
