package httpx

import (
	"net/http"
)

func Do(client *http.Client, request *http.Request) (*Response, error) {
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	return WrapResponse(resp), nil
}
