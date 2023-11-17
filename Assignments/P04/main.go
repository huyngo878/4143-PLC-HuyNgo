package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"path/filepath"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for i, url := range urls {
		filename := filepath.Join("images", fmt.Sprintf("image_%d.jpg", i))
		err := downloadImage(url, filename)
		if err != nil {
			fmt.Printf("Failed to download %s: %v\n", url, err)
		}
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for i, url := range urls {
		wg.Add(1)
		go func(url string, i int) {
			defer wg.Done()
			filename := filepath.Join("images", fmt.Sprintf("image_%d.jpg", i))
			err := downloadImage(url, filename)
			if err != nil {
				fmt.Printf("Failed to download %s: %v\n", url, err)
			}
		}(url, i)
	}
	wg.Wait()
}

func main() {
	urls := []string{
		"https://cdn.pixabay.com/photo/2023/11/08/20/11/mountains-8375693_1280.jpg",
		"https://cdn.pixabay.com/photo/2023/09/16/18/18/wallpaper-8257343_1280.png",
		"https://cdn.pixabay.com/photo/2016/06/02/02/33/triangles-1430105_1280.png",
		"https://cdn.pixabay.com/photo/2023/11/11/17/02/sunset-8381528_1280.jpg",
    "https://images.unsplash.com/photo-1700041829045-530ffd99f435?q=80&w=1587&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	// Create the directory if it does not exist
	dir := filepath.Dir(filename)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, 0755) // Using 0755 as the directory permission
		if err != nil {
			return err
		}
	}

	// Send HTTP GET request
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write response to file
	_, err = io.Copy(file, response.Body)
	return err
}
