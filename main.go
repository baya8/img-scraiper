package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	fmt.Println("Enter the site url")

	var site string
	fmt.Println("input url")
	fmt.Scan(&site)

	if isInvalid(site) {
		fmt.Println("有効なURLではありません")
		os.Exit(1)
	}

	execute(site)
}

func isInvalid(url string) bool {

	_, err := http.Get(url)
	return err != nil
}

func execute(input string) {

	var urls []*url.URL

	base, _ := url.Parse(input)

	doc, _ := goquery.NewDocument(input)
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {

		src, exists := s.Attr("src")
		if exists {
			srcs, _ := url.Parse(src)
			urls = append(urls, base.ResolveReference(srcs))
		}
	})

	for id, body := range urls {

		_, filename := path.Split(body.String())
		filepath := fmt.Sprintf("%d_%s", id, filename)
		fmt.Println(filepath)
		// exec.Command("wget", body.String()).Output()

		res, _ := http.Get(body.String())
		body, _ := ioutil.ReadAll(res.Body)
		file, _ := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY, 0666)

		file.Write(body)
		file.Close()
	}
}
