package pkg

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)


var templates = template.Must(template.ParseGlob("static/templates/*.html"))

type Images struct {
	ImagePaths []string
}

type IndexImage struct {
    Path string
}

type Videos struct {
    VideoPaths []string
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
    pageVariables := Images { ImagePaths: imgPaths }
    renderTemplate(w, "journal.html", pageVariables)
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "about.html", nil)
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
    err, allImages := getAllImages();
    if err != nil {
        http.Error(w, "Error getting all images", http.StatusInternalServerError)
        return
    }
    renderTemplate(w, "admin_page.html", Images { ImagePaths: allImages })
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
        fmt.Println("There's no index image")
        return "", nil
    }
    if len(matches) > 1 {
        fmt.Printf("Found more than one index image, returning first match")
        return matches[0], nil
    }
    return matches[0], nil
}

func getAllImages() (error, []string) {
    var images []string
    err := filepath.Walk("static/images", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if !info.IsDir() {
            images = append(images, path)
        }
        return nil
    })
    if err != nil {
        fmt.Println("Error while getting all images")
        return err, nil
    }
    return nil, images
}

func getAllVideos() (error, []string) {
    var videos []string
    err := filepath.Walk("static/videos", func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            videos = append(videos, path)
        }
        return nil
    })
    if err != nil {
        fmt.Println("Error while getting all videos")
        return err, nil
    }
    return nil, videos
}
