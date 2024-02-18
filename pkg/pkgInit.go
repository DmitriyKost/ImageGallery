package pkg

import (
	"net/http"
	"github.com/DmitriyKost/ImageGallery/pkg/database"
)

func init() {
    err := database.InitDatabase()
    if err != nil {
        panic(err)
    }
}

func InitRoutes() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/journal", JournalHandler)
    http.HandleFunc("/projects", ProjectsHandler)
    http.HandleFunc("/about", AboutHandler)
    http.HandleFunc("/login", Login)
    http.Handle("/admin", AuthMiddleWare(http.HandlerFunc(AdminHandler)))
    http.Handle("/upload", AuthMiddleWare(http.HandlerFunc(UploadHandler)))
    http.Handle("/replace_idx", AuthMiddleWare(http.HandlerFunc(ReplaceIndexImageHandler)))
    http.Handle("/delete_image", AuthMiddleWare(http.HandlerFunc(DeleteHandler)))
}
