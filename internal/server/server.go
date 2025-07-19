package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Loger      *log.Logger
	HttpServer http.Server
}

func CreateServer(loger *log.Logger) *Server {
	route := http.NewServeMux()
	route.HandleFunc("/", handlers.ReturnHTML)
	route.HandleFunc("/upload", handlers.UploadHTML)
	s := Server{
		Loger: loger,
		HttpServer: http.Server{
			Addr:         ":8080",
			Handler:      route,
			ErrorLog:     loger,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  15 * time.Second,
		},
	}
	return &s
}
