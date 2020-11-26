package net

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/adrg/xdg"
)

var ErrorEmptyResponse = errors.New("Empty Response")

// HTTPGetWithKey -- get contents from url, with cache and optional key
func HTTPGetWithKey(url string, keyName string, keyValue string) ([]byte, error) {
	URL := regexp.MustCompile("^https*://").ReplaceAllString(url, "")
	// cacheFile := "/home/box/.cache/book/" + URL

	cacheFile, err := xdg.CacheFile(Koanf.String("cache") + "/" + URL)
	if err != nil {
		return nil, err
	}

	//Logger.Debugf("GET: %v", URL)

	if _, err := os.Stat(cacheFile); err == nil {
		fileContent, readError := ioutil.ReadFile(cacheFile)
		if readError == nil && len(fileContent) <= 2 {
			return fileContent, ErrorEmptyResponse
		}

		return fileContent, readError
	}

	URL = url
	if len(keyName) > 0 {
		char := "?"
		if strings.Index(url, char) > 0 {
			char = "&"
		}
		URL = url + fmt.Sprintf(char+"%v=%v", keyName, keyValue)
	}
	resp, err := http.Get(URL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	htmlData, readErr := ioutil.ReadAll(resp.Body)

	if readErr != nil {
		return nil, err
	}
	cacheDir := filepath.Dir(cacheFile)
	err = os.MkdirAll(cacheDir, os.ModePerm)
	if err == nil {
		writeError := ioutil.WriteFile(cacheFile, htmlData, 0644)
		return htmlData, writeError
	}

	return nil, err
}

// HTTPGet -- get contents from url, with cache
func HTTPGet(url string) ([]byte, error) {
	return HTTPGetWithKey(url, "", "")
}
