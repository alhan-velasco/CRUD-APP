package server

import (
    "sync"
    "crud-app/shared"
)

var (
    users []shared.User
    mutex sync.Mutex
)

func AddUser(user shared.User) {
    mutex.Lock()
    defer mutex.Unlock()
    users = append(users, user)
}

func GetUsers() []shared.User {
    mutex.Lock()
    defer mutex.Unlock()
    return users
}

func DeleteUser(id int) {
    mutex.Lock()
    defer mutex.Unlock()
    for i, user := range users {
        if user.ID == id {
            users = append(users[:i], users[i+1:]...)
            break
        }
    }
}