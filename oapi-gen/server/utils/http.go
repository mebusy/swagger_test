package utils

import (
	"io/ioutil"
	"net/http"
)

func GetRequestBody(r *http.Request) ([]byte, error) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		return nil, err
	}
	return body, nil
}
