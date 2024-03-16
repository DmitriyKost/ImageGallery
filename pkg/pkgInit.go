package pkg

import (
	"net/http"

	"github.com/DmitriyKost/ImageGallery/pkg/database"
	"github.com/DmitriyKost/ImageGallery/pkg/structs"
)

type Image = structs.Image
type Video = structs.Video


func init() {
    err := database.InitDatabase()
    if err != nil {
        panic(err)
    }
}

func InitRoutes() {
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    http.HandleFunc("/", IndexHandler)
    http.HandleFunc("/edit_desc", EditDescHandler)
    http.HandleFunc("/journal", JournalHandler)
    http.HandleFunc("/videos", VideoJournalHandler)
    http.HandleFunc("/projects", ProjectsHandler)
    http.HandleFunc("/about", AboutHandler)
    http.HandleFunc("/login", Login)
    http.Handle("/admin", AuthMiddleWare(http.HandlerFunc(AdminHandler)))
    http.Handle("/upload", AuthMiddleWare(http.HandlerFunc(UploadHandler)))
    http.Handle("/replace_idx", AuthMiddleWare(http.HandlerFunc(ReplaceIndexImageHandler)))
    http.Handle("/delete", AuthMiddleWare(http.HandlerFunc(DeleteHandler)))
    // http.Handle("/edit_desc", AuthMiddleWare(http.HandlerFunc(EditDescHandler)))
}
