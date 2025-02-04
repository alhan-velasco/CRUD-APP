package replica

import (
    "encoding/json"
    "net/http"
    "crud-app/shared"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    users := shared.Users
    json.NewEncoder(w).Encode(users)
}