package replica

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func Start() {
    r := mux.NewRouter()
    r.HandleFunc("/users", GetUsersHandler).Methods("GET")

    log.Println("Servidor de réplica iniciado en :8081")
    log.Fatal(http.ListenAndServe(":8081", r))
}