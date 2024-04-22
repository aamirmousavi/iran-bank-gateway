package novinpal

import (
	"io"
	"net/http"
	"net/url"
)

func request(
	url string,
	formData url.Values,
) (
	*int,
	[]byte,
	error,
) {
	response, err := http.PostForm(url, formData)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	return &response.StatusCode, responseBody, nil
}
