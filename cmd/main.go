package main

import (
    "net/http"
    "github.com/DmitriyKost/ImageGallery/pkg"
)

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", pkg.IndexHandler)
    http.HandleFunc("/journal", pkg.JournalHandler)
    http.HandleFunc("/projects", pkg.ProjectsHandler)
    http.HandleFunc("/about", pkg.AboutHandler)
    http.ListenAndServe(":8080", nil)
}
