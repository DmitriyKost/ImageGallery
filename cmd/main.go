package main


import (
    "net/http"
    "github.com/DmitriyKost/ImageGallery/pkg"
)

func index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/upload", pkg.UploadHandler)
    http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)
}
