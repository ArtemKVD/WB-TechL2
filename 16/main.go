package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

func main() {
	url := flag.String("url", "", "")
	depth := flag.Int("depth", 1, "")
	output := flag.String("outfolder", "", "")
	flag.Parse()

	err := os.MkdirAll(*output, 0755)
	if err != nil {
		log.Println("error create directory")
	}

	err = download(*url, *depth, *output)
	if err != nil {
		log.Println("error download: ", err)
	}

}

func download(startURL string, Depth int, outputDir string) error {
	baseURL, err := url.Parse(startURL)
	if err != nil {
		return err
	}
	domain := baseURL.Hostname()

	c := colly.NewCollector(
		colly.AllowedDomains(domain),
		colly.MaxDepth(Depth),
		colly.Async(true),
	)

	c.OnHTML("html", func(e *colly.HTMLElement) {
		saveContent(e.Request.URL.String(), []byte(e.Response.Body), "text/html", outputDir)
	})

	c.OnHTML("link[href], script[src], img[src], source[src]", func(e *colly.HTMLElement) {
		var resourceURL string

		switch {
		case e.Attr("href") != "" && strings.Contains(e.Attr("rel"), "stylesheet"):
			resourceURL = e.Attr("href")
		case e.Attr("src") != "":
			resourceURL = e.Attr("src")
		default:
			return
		}

		URL := e.Request.AbsoluteURL(resourceURL)
		if URL != "" {
			downloadResource(URL, outputDir)
		}
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link != "" {
			URL := e.Request.AbsoluteURL(link)
			if URL != "" {
				e.Request.Visit(URL)
			}
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("downloading error: ", err)
	})

	err = c.Visit(startURL)
	if err != nil {
		return err
	}

	c.Wait()

	return nil
}

func downloadResource(url string, outputDir string) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("creating request error: ", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("downloading error: ", err)
		return
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("reading error", err)
		return
	}

	contentType := resp.Header.Get("Content-Type")
	saveContent(url, content, contentType, outputDir)
}

func saveContent(URL string, content []byte, contentType string, outputDir string) {
	parsedURL, err := url.Parse(URL)
	if err != nil {
		log.Println("parsing URL error: ", err)
		return
	}

	domainDir := filepath.Join(outputDir, parsedURL.Hostname())
	filename := getFilename(parsedURL, contentType)
	filePath := filepath.Join(domainDir, filename)

	err = os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		log.Println("creating directory error: ", err)
		return
	}

	err = os.WriteFile(filePath, content, 0644)
	if err != nil {
		log.Println("saving error:", filePath, err)
		return
	}
}

func getFilename(parsedURL *url.URL, contentType string) string {
	path := parsedURL.Path
	if path == "" || path == "/" {
		return "index.html"
	}

	if strings.HasPrefix(path, "/") {
		path = path[1:]
	}

	if strings.HasSuffix(path, "/") {
		return path + "index.html"
	}

	if !strings.Contains(filepath.Base(path), ".") {
		switch {
		case strings.Contains(contentType, "text/html"):
			return path + ".html"
		case strings.Contains(contentType, "text/css"):
			return path + ".css"
		case strings.Contains(contentType, "javascript"):
			return path + ".js"
		case strings.Contains(contentType, "image/jpeg"):
			return path + ".jpg"
		}
	}

	return path
}
