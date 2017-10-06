package utils

import (
	"fmt"
	"time"
	"github.com/cavaliercoder/grab"
	"net/http"
	"io/ioutil"
	"errors"
)

func DownloadURL(url string, file string) bool {
	fmt.Printf("Downloading %s...\n", url)
	client := grab.NewClient()

	req, _ := grab.NewRequest(file, url)

	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)
	if(resp.Size == 0){
		return false
	}

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

func DownloadString(url string) (string, error) {
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		return  "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return  "", err2
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	}

	return  "", errors.New("Failed to download file")
}
