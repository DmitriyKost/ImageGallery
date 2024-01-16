package main


import (
    "net/http"
    "github.com/DmitriyKost/ImageGallery/pkg"
)

func index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "/Users/dmitriyk/.some_code/my_pet_projects/ImageGallery/static/index.html")
}

func main() {
    http.HandleFunc("/upload", pkg.UploadHandler)
    http.HandleFunc("/", index)
    http.ListenAndServe(":8080", nil)
}
