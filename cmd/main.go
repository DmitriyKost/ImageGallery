package main

import (
    "net/http"
    "github.com/joho/godotenv"
    "github.com/DmitriyKost/ImageGallery/pkg"
)
func init() { 
    if err := godotenv.Load("creds.env"); err != nil {
        panic("Error loading credentials")
    }
    pkg.InitEnv()
}

func main() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", pkg.IndexHandler)
    http.HandleFunc("/journal", pkg.JournalHandler)
    http.HandleFunc("/projects", pkg.ProjectsHandler)
    http.HandleFunc("/about", pkg.AboutHandler)
    http.HandleFunc("/login", pkg.Login)
    http.Handle("/admin", pkg.AuthMiddleWare(http.HandlerFunc(pkg.AdminHandler)))
    http.ListenAndServe(":8080", nil)
}
