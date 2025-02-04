package main

import (
    "crud-app/server"
    "crud-app/replica"
)

func main() {
    go server.Start()
    replica.Start()
}