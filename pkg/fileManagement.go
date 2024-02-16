package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
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

    // TODO handle videos
    newFile, err := os.Create("static/images/journal/" + handler.Filename)
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

    response := Response{Message: "Image uploaded successfully"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
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
        http.Error(w, "Error replacing index_image", http.StatusInternalServerError)
    }
    if len(matches) > 0 {
        os.Remove(matches[0])
    }

    fileExt := strings.Split(handler.Filename, ".")[1]
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

    response := Response{Message: "Image replaced successfully"}
    jsonResponse, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
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

    err = os.Remove(path)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error while deleting: %s", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message: Successfully deleted"}`))
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
