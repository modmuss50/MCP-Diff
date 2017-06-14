package utils

import (
	"fmt"
	"time"
	"github.com/cavaliercoder/grab"
)

func DownloadURL(url string, file string) bool {
	fmt.Printf("Downloading %s...\n", url)
	client := grab.NewClient()

	req, _ := grab.NewRequest(file, url)

	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			fmt.Printf("  transferred %v / %v bytes (%.2f%%)\n",
				resp.BytesComplete(),
				resp.Size,
				100*resp.Progress())

		case <-resp.Done:
			break Loop
		}
	}

	if err := resp.Err(); err != nil {
		return false
	}

	fmt.Printf("Download saved to ./%v \n", resp.Filename)

	return true
}
