package main

import (
    "github.com/Fodi999/cssgenserverdum/util"
    "log"
    "net/http"
    
)

func main() {
    // Создание необходимых файлов при старте сервера
    util.GenerateCSS()
    util.CreateDefaultHTMLFile()
    util.CreateDefaultJSFile()

    // Роутинг
    http.HandleFunc("/generate-css", generateCSSHandler)
    http.HandleFunc("/create-site", createSiteHandler)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.Handle("/templates/", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))

    // Запуск сервера
    log.Println("Starting server on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateCSSHandler(w http.ResponseWriter, r *http.Request) {
    util.GenerateCSS()
    w.Write([]byte("CSS generated successfully."))
}

func createSiteHandler(w http.ResponseWriter, r *http.Request) {
    siteName := r.URL.Query().Get("name")
    if siteName == "" {
        http.Error(w, "Missing 'name' parameter", http.StatusBadRequest)
        return
    }

    err := util.CreateSite(siteName)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Site created successfully."))
}

