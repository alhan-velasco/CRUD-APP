package server

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func Start() {
    r := mux.NewRouter()
    r.HandleFunc("/users", AddUserHandler).Methods("POST")
    r.HandleFunc("/users", GetUsersHandler).Methods("GET")
    r.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")
    r.HandleFunc("/shortpoll", ShortPollHandler).Methods("GET")

    log.Println("Server principal iniciado en :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}