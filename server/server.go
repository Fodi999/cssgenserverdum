package server

import (
    
    "log"
    "net/http"
    "os"
    "path/filepath"

    "github.com/ваш_пользователь/github-репозиторий/util"
)

// cssHandler обрабатывает запросы и запускает генерацию CSS
func cssHandler(w http.ResponseWriter, r *http.Request) {
    util.GenerateCSS()

    cssFile, err := os.ReadFile(filepath.Join("static", "css", "style.css"))
    if err != nil {
        http.Error(w, "Error reading CSS file", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/css")
    w.Write(cssFile)
}

// StartServer запускает HTTP сервер на указанном порту
func StartServer(port string) {
    http.HandleFunc("/generate", cssHandler)
    log.Printf("Starting server on %s\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}
