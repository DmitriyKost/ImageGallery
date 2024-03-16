package database

import (
	"database/sql"
	"os"
	"strings"
	"sync"

	_ "github.com/DmitriyKost/ImageGallery/env"
	"github.com/DmitriyKost/ImageGallery/pkg/structs"
	_ "github.com/mattn/go-sqlite3"
)

var DataBasePath = os.Getenv("DB_PATH")

var Database *sql.DB
var mutex sync.Mutex

type Image = structs.Image
type Video = structs.Video
type Item = structs.Item

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

    err, imagePaths := getAllImagePaths()
    if err != nil {
        return err
    }
    err, videoPaths := getAllVideoPaths()
    if err != nil {
        return err
    }

    query := "INSERT INTO images (path) VALUES (?)"
    for _, path := range imagePaths {
        row := Database.QueryRow("SELECT id FROM images WHERE path = ?", path)
        var id int
        err := row.Scan(&id)
        if err != nil && err != sql.ErrNoRows {
            return err
        } else if err == sql.ErrNoRows {
            _, err := Database.Exec(query, path)
            if err != nil {
                return err
            }
        }
    }

    query = "INSERT INTO videos (path) VALUES (?)"
    for _, path := range videoPaths {
        row := Database.QueryRow("SELECT id FROM videos WHERE path = ?", path)
        var id int
        err := row.Scan(&id)
        if err != nil && err != sql.ErrNoRows {
            return err
        } else if err == sql.ErrNoRows {
            _, err := Database.Exec(query, path)
            if err != nil {
                return err
            }
        }
    }
    return nil
}

func GetAll() ([]Item, error) {
    rows, err := Database.Query("SELECT id, path, description FROM images UNION SELECT id, path, description FROM videos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var items []Item
	for rows.Next() {
        var item Item
		err := rows.Scan(&item.Id, &item.Path, &item.Description)
		if err != nil {
			return nil, err
		}
        items = append(items, item)
	}
    return items, nil

}

func GetAllImages() ([]Image, error) {
	rows, err := Database.Query("SELECT id, path, description FROM images;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var images []Image
	for rows.Next() {
        var image Image
		err := rows.Scan(&image.Id, &image.Path, &image.Description)
		if err != nil {
			return nil, err
		}
        images = append(images, image)
	}
    return images, nil
}

func GetAllVideos() ([]Video, error) {
	rows, err := Database.Query("SELECT id, path, description FROM videos;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

    var videos []Video
	for rows.Next() {
        var video Video
		err := rows.Scan(&video.Id, &video.Path, &video.Description)
		if err != nil {
			return nil, err
		}
        videos = append(videos, video)
	}
    return videos, nil
}

func GetJournalImages() (error, []Image) {
    var images []Image
    paths, err := getJournalImagePaths()
    if err != nil {
        return err, nil
    }
    for _, path := range paths {
        var image Image
        row := Database.QueryRow("SELECT id, path, description FROM images WHERE path = ?", path)
        err := row.Scan(&image.Id, &image.Path, &image.Description)
        if err != nil {
            return err, nil
        }
        images = append(images, image)
    }
    return nil, images
}

func GetJournalVideos() (error, []Video) {
    var videos []Video
    paths, err := getJournalVideoPaths()
    if err != nil {
        return err, nil
    }
    for _, path := range paths {
        var video Video
        row := Database.QueryRow("SELECT id, path, description FROM videos WHERE path = ?", path)
        err := row.Scan(&video.Id, &video.Path, &video.Description)
        if err != nil {
            return err, nil
        }
        videos = append(videos, video)
    }
    return nil, videos
}

func GetIndexImage() (error, Image) {
    var idxImage Image
    path, err := getIndexImagePath()
    if err != nil {
        return err, idxImage
    }
    row := Database.QueryRow("SELECT id, path, description FROM images WHERE path = ?", path)
    err = row.Scan(&idxImage.Id, &idxImage.Path, &idxImage.Description)
    if err != nil {
        return err, idxImage
    }
    return nil, idxImage
}

func DeleteFromDB(path string) error {
    mutex.Lock()
    defer mutex.Unlock()
    toDelete := strings.Split(path, "/")[1] // path usually looks like "static/videos" or "static/images"
    query := "DELETE FROM " + toDelete + " WHERE path = ?"
    _, err := Database.Exec(query, path)
    if err != nil {
        return err
    }
    return nil
}

func InsertItem(path string) error {
    toInsert:= strings.Split(path, "/")[1] // path usually looks like "static/videos" or "static/images"
    _, err := Database.Exec("INSERT INTO " + toInsert + " (path) VALUES (?)", path)
    if err != nil {
        return err
    }
    return nil
}

func EditDescription(path string, description string) error {
    mutex.Lock()
    defer mutex.Unlock()

    toEdit := strings.Split(path, "/")[1] // path usually looks like "static/videos" or "static/images"
    query := "UPDATE " + toEdit + " SET description = ? WHERE path = ?"
    _, err := Database.Exec(query, description, path);
    if err != nil {
        return err
    }
    return nil
}
