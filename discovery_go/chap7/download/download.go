package main

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

// download downloads url and returns the contents and error.
func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

// urlToFilename returns the filename part from the rawurl
func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil
}

// writeZip writes a zip archive file.
func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}

func main() {
	var urls = []string{
		"http://image.com/img/home/hero-2.jpg",
		"http://image.com/img/home/hero-5.jpg",
		"http://image.com/img/home/hero-3.jpg",
	}

	// 고루틴 기다리기
	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()

	filenames, err := filepath.Glob("*.jpg")
	if err != nil {
		log.Fatal(err)
	}
	err = writeZip("images.zip", filenames)
	if err != nil {
		log.Fatal(err)
	}
}
