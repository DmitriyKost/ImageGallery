package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	"github.com/DmitriyKost/ImageGallery/pkg/database"
)

type Response struct {
    Message string `json:"message"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(100 << 20) // 100 MB limit
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("item")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    ft, err := fileType(handler.Filename)
    if err != nil {
        http.Error(w, "Unknown file extension", http.StatusBadRequest)
        return
    }

    err, freq := fileFreq(handler.Filename)
    if err != nil {
        fmt.Println("Error walking directories")
    }

    prefix := fmt.Sprintf("(%d)", freq)

    newFile, err := os.Create("static/" + ft + "/journal/" + prefix + handler.Filename)
    if err != nil {
        http.Error(w, "Error creating file", http.StatusInternalServerError)
        return
    }
    defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        http.Error(w, "Error copying file", http.StatusInternalServerError)
        return
    }

    err = database.InsertItem("static/" + ft + "/journal/" + prefix + handler.Filename)
    if err != nil {
        http.Error(w, "Database error while inserting file", http.StatusInternalServerError)
        return
    }

    response := Response{Message: "File uploaded successfully"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func ReplaceIndexImageHandler(w http.ResponseWriter, r *http.Request) {
    err := r.ParseMultipartForm(10 << 20) // 10 MB limit
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    file, handler, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    matches, err := filepath.Glob("static/images/index_image.*")
    if err != nil {
        fmt.Println("No image to replace")
    }
    if len(matches) > 0 {
        database.DeleteFromDB(matches[0])
        os.Remove(matches[0])
    }


    fileParts := strings.Split(handler.Filename, ".")
    fileExt := fileParts[len(fileParts)-1]
    newFileName := "index_image." + fileExt

    newFile, err := os.Create("static/images/" + newFileName)
    if err != nil {
        http.Error(w, "Error creating file", http.StatusInternalServerError)
        return
    }
    defer newFile.Close()

    _, err = io.Copy(newFile, file)
    if err != nil {
        http.Error(w, "Error copying file", http.StatusInternalServerError)
        return
    }

    database.InsertItem("static/images/" + newFileName)

    response := Response{Message: "Image replaced successfully"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    body, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }

    path := string(body)
    if !isPathSecure(path) {
        http.Error(w, "Path is invalid", http.StatusBadRequest)
        return
    }

    database.DeleteFromDB(path)

    err = os.Remove(path)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error while deleting: %s", err), http.StatusInternalServerError)
        return
    }

    response := Response{Message: "Successfully deleleted"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func EditDescHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPatch {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    if err := r.ParseForm(); err != nil {
        http.Error(w, "Error parsing form data", http.StatusBadRequest)
        return
    }

    desc, path := r.Form.Get("desc"), r.Form.Get("path")

    if !isPathSecure(path) {
        http.Error(w, "Path is invalid", http.StatusBadRequest)
        return
    }

    database.EditDescription(path, desc)

    response := Response{Message: "Successfully updated"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}


func isPathSecure(path string) bool {
    formats := []string{ "images", "videos" }

    if path[0] == '/' {
        return false
    }
    if strings.Contains(path, "..") {
        return false
    }
    splitted := strings.Split(path, "/")
    if splitted[0] != "static" || !slices.Contains(formats, splitted[1]) {
        return false
    }
    return true
}

// Checks for the supported image/video format and determines filetype, or returns an error if unknown file format
func fileType(filename string) (string, error) {
	lowercaseFilename := strings.ToLower(filename)

	if strings.HasSuffix(lowercaseFilename, ".mp4") ||
		strings.HasSuffix(lowercaseFilename, ".avi") ||
		strings.HasSuffix(lowercaseFilename, ".mov") {
        return "videos", nil
	} else if strings.HasSuffix(lowercaseFilename, ".jpg") ||
		strings.HasSuffix(lowercaseFilename, ".jpeg") ||
		strings.HasSuffix(lowercaseFilename, ".png") ||
		strings.HasSuffix(lowercaseFilename, ".gif") {
        return "images", nil
	} else {
        return "", fmt.Errorf("Unknown file extension")
	}
}

func fileFreq(filename string) (error, int) {
    escapedFilename := regexp.QuoteMeta(filename)
	pattern1 := escapedFilename
	pattern2 := fmt.Sprintf(`\(\d+\)%s`, escapedFilename)
    matches := 0
    matcher := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		matched, err := regexp.MatchString(pattern1, info.Name())
		if err != nil {
			return err
		}

		matched2, err := regexp.MatchString(pattern2, info.Name())
		if err != nil {
			return err
		}

		if matched || matched2 {
            matches++
		}

		return nil
	}
    err := filepath.Walk("static/", matcher)
	if err != nil {
		return err, -1
	}
    return nil, matches
}
