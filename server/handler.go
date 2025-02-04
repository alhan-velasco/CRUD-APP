package server

import (
    "encoding/json"
    "net/http"
    "strconv"
    "crud-app/shared"
    "github.com/gorilla/mux"
)

func AddUserHandler(w http.ResponseWriter, r *http.Request) {
    var user shared.User
    json.NewDecoder(r.Body).Decode(&user)
    AddUser(user)
    w.WriteHeader(http.StatusCreated)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := GetUsers()
    json.NewEncoder(w).Encode(users)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    DeleteUser(id)
    w.WriteHeader(http.StatusNoContent)
}