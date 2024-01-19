package pkg

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)


var templates = template.Must(template.ParseGlob("static/templates/*.html"))

type JournalImages struct {
	ImagePaths []string
}

type IndexImage struct {
    Path string
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
    imgPath, err := getIndexImage()
    if err != nil {
        http.Error(w, "Error reading index image", http.StatusInternalServerError)
        return
    }
    idxImage := IndexImage { imgPath }
    renderTemplate(w, "index.html", idxImage)
}

func ProjectsHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "projects.html", nil)
}

func JournalHandler(w http.ResponseWriter, r *http.Request) {
    imgPaths, err := getJournalImages()
    if err != nil {
        http.Error(w, "Error reading journal image directory", http.StatusInternalServerError)
        return
    }
    pageVariables := JournalImages { ImagePaths: imgPaths }
    renderTemplate(w, "journal.html", pageVariables)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.html", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data any) {
    err := templates.ExecuteTemplate(w, tmpl, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}

func getJournalImages() ([]string, error) {
	var paths []string

	err := filepath.Walk("static/images/journal",
        func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return paths, nil
}

func getIndexImage() (string, error) {
    matches, err := filepath.Glob("static/images/index_image.*")
    if err != nil {
        return "", err
    }
    if len(matches) < 1 {
        fmt.Printf("There's no index image")
        return "", nil
    }
    if len(matches) > 1 {
        fmt.Printf("Found more than one index image, returning first match")
        return matches[0], nil
    }
    return matches[0], nil
}
