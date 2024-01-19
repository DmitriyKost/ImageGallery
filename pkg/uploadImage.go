package pkg


import (
    "net/http"
    "io"
    "os"
    "encoding/json"
)

type Response struct {
    Message string `json:"message"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
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

    newFile, err := os.Create("static/images/" + handler.Filename)
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
