package main

import (
	"log"
	"net/http"

	"github.com/DmitriyKost/ImageGallery/pkg"
)


func main() {
    pkg.InitRoutes();
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Println("Error listening on :8080")
    }
}
