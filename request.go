package ItalyPassportAlert

import (
	"github.com/corpix/uarand"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	*http.Client
	*http.Request
}

func CreateRequest(location string) (Request, error) {
	request, err := http.NewRequest(http.MethodGet, "https://www.passaportonline.poliziadistato.it/CittadinoAction.do", nil)
	if err != nil {
		return Request{nil, nil}, err
	}

	// Imitate a browser
	request.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	request.Header.Set("Accept-Encoding", "gzip, deflate, br")
	request.Header.Set("Accept-Language", "it-IT,it;q=0.9")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Host", "gzip, deflate, br")
	request.Header.Set("Cache-Control", "max-age=0")
	request.Header.Set("Upgrade-Insecure-Requests", "1")

	query := request.URL.Query()
	query.Add("codop", "resultRicercaRegistiProvincia")
	query.Add("provincia", location)
	request.URL.RawQuery = query.Encode()

	return Request{&http.Client{}, request}, nil
}

func (r *Request) getBodyContent() ([]byte, error) {
	// Random user agent each time
	r.Header.Set("User-Agent", uarand.GetRandom())
	response, err := r.Client.Do(r.Request)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(response.Body)
}

func (r *Request) CheckAvailability() bool {
	body, err := r.getBodyContent()
	if err != nil {
		return false
	}
	return strings.Contains(string(body), "<td headers=\"disponibilita\">Si</td>")
}
