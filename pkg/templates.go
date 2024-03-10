package pkg

import (
	"html/template"
	"net/http"

	"github.com/DmitriyKost/ImageGallery/pkg/database"
)


var templates = template.Must(template.ParseGlob("static/templates/*.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    err, idxImage := database.GetIndexImage()
    if err != nil {
        renderTemplate(w, "index.html", nil)
        return
    }
    renderTemplate(w, "index.html", idxImage)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "projects.html", nil)
}

func JournalHandler(w http.ResponseWriter, r *http.Request) {
    err, JournalImages := database.GetJournalImages()
    if err != nil {
        http.Error(w, "Error reading journal image directory", http.StatusInternalServerError)
        return
    }
    renderTemplate(w, "journal.html", JournalImages)
}
func VideoJournalHandler(w http.ResponseWriter, r *http.Request) {
    err, JournalVideos := database.GetJournalVideos()
    if err != nil {
        http.Error(w, "Error reading journal video directory", http.StatusInternalServerError)
        return
    }
    renderTemplate(w, "videos.html", JournalVideos)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.html", nil)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
    allFiles, err := database.GetAll();
    if err != nil {
        http.Error(w, "Error getting all files", http.StatusInternalServerError)
        return
    }
    renderTemplate(w, "admin_page.html", allFiles)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
