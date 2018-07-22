package userscloud

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func DownloadFile(pathToSave string, url string) error {
	err := downloadFile(pathToSave, getDLLinkInUserscloud(url))
	if err != nil {
		return err
	}
	return nil
}

func downloadFile(pathToSave string, url string) error {
	log.Println("Crawler:", url)

	// Create the file
	out, err := os.Create(pathToSave)
	if err != nil {
		log.Println("Error!:", err)
		return err
	}
	log.Println("Create file success. => ", pathToSave)
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Println("Error!:", err)
		return err
	}
	defer resp.Body.Close()
	log.Println("Init download ...")

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Println("Error!:", err)
		return err
	}
	log.Println("Download success!")
	return nil
}

func getDLLinkInUserscloud(url string) string {
	// Get key
	log.Println("Crawler:", url)
	key := strings.Split(url, "userscloud.com/")
	payload := strings.NewReader("op=download2&id=" + key[1])
	log.Println("key:", key[1])

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	link := ""
	doc.Find(".btn-icon-stacked").Each(func(i int, selection *goquery.Selection) {
		link, _ = selection.Attr("href")
	})
	return link
}
