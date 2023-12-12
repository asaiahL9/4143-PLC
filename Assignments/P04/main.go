package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Sequential version of the image downloader.
// downloadImagesSequential sequentially downloads images from the provided list of URLs.
func downloadImagesSequential(urls []string) {
	// Iterate through each URL in the list.
	for _, url := range urls {
		// Generate a filename based on the URL.
		filename := generateFilename(url)

		// Download the image from the current URL and save it with the generated filename.
		err := downloadImage(url, filename)

		// Check if there was an error during the download.
		if err != nil {
			// Print an error message, including the URL and the specific error.
			fmt.Printf("Error downloading image from %s: %v\n", url, err)
		}
	}
}


// generateFilename generates a unique filename for the given URL based on the current timestamp.
func generateFilename(url string) string {
	// Extract the file extension from the URL.
	fileExt := filepath.Ext(url)

	// Generate a filename using the current Unix timestamp in nanoseconds
	return fmt.Sprintf("image_%d%s", time.Now().UnixNano(), fileExt)
}


// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			filename := generateFilename(url)
			err := downloadImage(url, filename)
			if err != nil {
				fmt.Printf("Error downloading image from %s: %v\n", url, err)
			}
		}(url)
	}

	wg.Wait()
}

func main() {
	urls := []string{
		"https://img.freepik.com/free-photo/medium-shot-man-wearing-vr-glasses_23-2150394443.jpg?w=900&t=st=1701962053~exp=1701962653~hmac=fd59704b81f45683cc7faa59dd43d13e0ae168291697af259e175b52a3b9dfc5",
		"https://img.freepik.com/free-photo/tall-lighthouse-north-sea-cloudy-sky_181624-49637.jpg?w=900&t=st=1701962092~exp=1701962692~hmac=5b97690d67a655cfee49c80ab2be3bd28c34f18f38d79647e895520518790926",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		// Add more image URLs
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
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request failed with status code %d", response.StatusCode)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Downloaded %s and saved as %s\n", url, filename)
	return nil
}
