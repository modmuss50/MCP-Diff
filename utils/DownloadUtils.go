package utils

import (
	"fmt"
	"github.com/cavaliercoder/grab"
	"os"
	"time"
)

func DownloadURL(url string, file string) bool {
	fmt.Printf("Downloading %s...\n", url)
	respch, err := grab.GetAsync(file, url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", url, err)
		return false
	}

	fmt.Printf("Initializing download...\n")
	resp := <-respch
	if resp.Size == 0 {
		fmt.Printf("Response size = 0")
		return false
	}

	for !resp.IsComplete() {
		fmt.Printf("\033[1AProgress %d / %d bytes (%d%%)\033[K\n", resp.BytesTransferred(), resp.Size, int(100*resp.Progress()))
		time.Sleep(200 * time.Millisecond)
	}

	fmt.Printf("\033[1A\033[K")

	if resp.Error != nil {
		fmt.Fprintf(os.Stderr, "Error downloading %s: %v\n", url, resp.Error)
		return false
	}

	fmt.Printf("Successfully downloaded to ./%s\n", resp.Filename)
	return true
}
