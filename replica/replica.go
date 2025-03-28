package replica

import (
    "encoding/json"
    "log"
    "net/http"
    "sync"
    "time"
    "crud-app/shared"
    "github.com/gorilla/mux"
)

var (
    lastUserCount int
    mu sync.Mutex
)

func LongPollHandler(w http.ResponseWriter, r *http.Request) {
    flusher, ok := w.(http.Flusher)
    if !ok {
        http.Error(w, "Streaming no soportado", http.StatusInternalServerError)
        return
    }

    for {
        mu.Lock()
        currentCount := len(shared.Users)
        mu.Unlock()

        if currentCount != lastUserCount {
            mu.Lock()
            lastUserCount = currentCount
            mu.Unlock()

            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(shared.Users)
            flusher.Flush()
            return
        }
        
        time.Sleep(2 * time.Second) 
    }
}

func Start() {
    r := mux.NewRouter()
    r.HandleFunc("/users", GetUsersHandler).Methods("GET")
    r.HandleFunc("/longpoll", LongPollHandler).Methods("GET")

    log.Println("Servidor de réplica iniciado en :8081")
    log.Fatal(http.ListenAndServe(":8081", r))
}
