package server

import (
	"ingot/config"
	"net/http"
	"time"
)

// HTTPServer 实例
func HTTPServer(handler http.Handler, conf config.Server) *http.Server {
	return &http.Server{
		Addr:         conf.Address,
		Handler:      handler,
		ReadTimeout:  conf.ReadTimeout * time.Second,
		WriteTimeout: conf.WriteTimeout * time.Second,
	}
}
