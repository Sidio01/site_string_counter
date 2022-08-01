package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHtmlPage(webPage string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s", webPage))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
