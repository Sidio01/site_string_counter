package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHtmlPage(webPage string) (string, string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s", webPage))
	if err != nil {
		return "", "404 Not Found", err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		return "", resp.Status, errors.New("status code not equal \"200 OK\"")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", resp.Status, err
	}
	return string(body), resp.Status, nil
}
