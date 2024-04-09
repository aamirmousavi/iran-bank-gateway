package novinpal

import (
	"bytes"
	"io"
	"net/http"
)

func request(
	method, url string,
	body []byte,
) (
	*int,
	[]byte,
	error,
) {
	request, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}
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
