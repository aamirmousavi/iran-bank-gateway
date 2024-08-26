package sep

import (
	"bytes"
	"io"
	"net/http"
)

func request(
	url string,
	body []byte,
) (*int, []byte, error) {
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer response.Body.Close()
	reponseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, nil, err
	}
	// fmt.Printf("response status code: %d\tresponse body: %s\n", response.StatusCode, reponseBody)
	return &response.StatusCode, reponseBody, nil
}
