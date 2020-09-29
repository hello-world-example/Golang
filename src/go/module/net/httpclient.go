package net

import (
	"net/http"
	"io/ioutil"
	"log"
)

func GetHtml(url string) (string, error) {
	response, err := http.Get(url)

	if nil != err {
		return "", err
	}

	defer response.Body.Close()

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Println(err)
		return "", err
	}

	return string(bytes), nil
}
