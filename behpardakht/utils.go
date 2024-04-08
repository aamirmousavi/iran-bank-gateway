package behpardakht

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
)

type soapRoot struct {
	XMLName xml.Name `xml:"x:Envelope"`
	X       string   `xml:"xmlns:x,attr"`
	Ns1     string   `xml:"xmlns:ns1,attr"`
	Body    soapBody
}

func (r *soapRoot) Marshal() ([]byte, error) {
	return xml.MarshalIndent(r, "", "  ")
}

type soapBody struct {
	XMLName xml.Name `xml:"x:Body"`
	Request interface{}
}

func newSoapRoot() *soapRoot {
	return &soapRoot{
		X:   "http://schemas.xmlsoap.org/soap/envelope/",
		Ns1: "http://interfaces.core.sw.bps.com/",
	}
}

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
	request.Header.Set("Content-Type", "text/xml")
	request.Header.Set("charset", "utf-8")
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
