package novinpal

import (
	"bytes"
	"io"
	"net/http"
)

func request(
	method, url string,
	contentType string,
	payload *bytes.Buffer,
) (
	*int,
	[]byte,
	error,
) {
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", contentType)
	response, err := http.DefaultClient.Do(request)
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
