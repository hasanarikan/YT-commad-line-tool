package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", os.Args[1], nil)
	errorcheck(err)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.138 Safari/537.36")
	resp, err := client.Do(req)
	errorcheck(err)
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	errorcheck(err)
	bodyString := string(bodyBytes)

	regex_title := regexp.MustCompile(`,"videoDetails":{"playerOverlayVideoDetailsRenderer":{"title":{"simpleText":"(.*?)"},`)
	title := regex_title.FindStringSubmatch(bodyString)
	regex_description := regexp.MustCompile(`"description":{"simpleText":"(.*?)"},`)
	description := regex_description.FindStringSubmatch(bodyString)

	regex_video_id := regexp.MustCompile(`<link rel="shortlinkUrl" href="https://youtu.be/(.*?)">`)
	video_id := regex_video_id.FindStringSubmatch(bodyString)
	// save output to json file
	f, err := os.Create(video_id[1] + ".json")
	errorcheck(err)
	defer f.Close()
	f.WriteString(`{"title":"` + title[1] + `","description":"` + description[1] + `"}`)

}

func errorcheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
