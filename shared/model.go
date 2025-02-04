package shared

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
}

var Users []User