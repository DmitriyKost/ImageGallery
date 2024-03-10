package database

import (
	"fmt"
	"os"
	"path/filepath"
)

func getJournalImagePaths() ([]string, error) {
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

func getJournalVideoPaths() ([]string, error) {
	var paths []string

	err := filepath.Walk("static/videos/journal",
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

func getIndexImagePath() (string, error) {
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

func getAllImagePaths() (error, []string) {
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

func getAllVideoPaths() (error, []string) {
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
