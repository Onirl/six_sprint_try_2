package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	loger := log.New(os.Stdout, "INFO:", log.Ldate|log.Ltime) //Создали логер
	myServer := server.CreateServer(loger)
	loger.Fatal(myServer.HttpServer.ListenAndServe())
}
