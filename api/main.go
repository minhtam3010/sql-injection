package main

import (
	"log"
	"net/http"

	"github.com/minhtam3010/sql-injection/handler"
)

func init() {
	log.Println("Server running at port 5500")
}
func main() {
	// Create a new router
	mux := http.NewServeMux()
	handler := handler.NewHandler()

	mux.HandleFunc("/", http.FileServer(http.Dir("../assets")).ServeHTTP)
	mux.HandleFunc("/create", handler.CreateUser)
	mux.HandleFunc("/login", handler.Login)

	err := http.ListenAndServe(":5500", mux)
	if err != nil {
		panic(err.Error())
	}
}
