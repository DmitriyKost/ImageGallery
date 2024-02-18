package database

import (
	"database/sql"
	"os"
	"strings"
	"sync"
    _ "github.com/DmitriyKost/ImageGallery/env"
	_ "github.com/mattn/go-sqlite3"
)

var DataBasePath = os.Getenv("DB_PATH")

var Database *sql.DB
var mutex sync.Mutex

type Image struct {
    Id int
    Name string
    Path string
    Description string
}

type Video struct {
    Id int
    Name string
    Path string
    Description string
}

func InitDatabase() error {
	if _, err := os.Stat(DataBasePath); os.IsNotExist(err) {
		os.Create(DataBasePath)
	}
    var err error
	Database, err = sql.Open("sqlite3", DataBasePath)
	data, err := os.ReadFile("./config/migrations.sql")
	if err != nil {
		return err
	}
	for _, query := range strings.Split(string(data), "query_separator") {
		_, err := Database.Exec(query)
		if err != nil {
			return err
		}
	}
    return nil
}

func GetAllImages() ([]Image, error) {
	rows, err := Database.Query("SELECT id, name, path, description FROM images;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var images []Image
	for rows.Next() {
        var image Image
		err := rows.Scan(&image.Id, &image.Name, &image.Path, &image.Description)
		if err != nil {
			return nil, err
		}
        images = append(images, image)
	}
    return images, nil
}

func GetAllVideos() ([]Video, error) {
	rows, err := Database.Query("SELECT id, name, path, description FROM videos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var videos []Video
	for rows.Next() {
        var video Video
		err := rows.Scan(&video.Id, &video.Name, &video.Path, &video.Description)
		if err != nil {
			return nil, err
		}
        videos = append(videos, video)
	}
    return videos, nil
}

func Delete(path string) error {
    mutex.Lock()
    defer mutex.Unlock()
    toDelete := strings.Split(path, "/")[1] // path usually looks like "static/videos" or "static/images"
    _, err := Database.Exec("DELETE FROM ? WHERE path = ?", toDelete, path)
    if err != nil {
        return err
    }
    return nil
}
