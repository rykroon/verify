package httpx

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type Request struct {
	*http.Request
}

func NewRequest(method, url string, body io.Reader) (*Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return &Request{req}, nil
}

func (r *Request) SetHeader(k, v string) *Request {
	r.Header.Set(k, v)
	return r
}

func (r *Request) SetContentType(contentType string) *Request {
	r.SetHeader("Content-Type", contentType)
	return r
}

func (r *Request) SetAuth(scheme, credential string) *Request {
	r.SetHeader("Authorization", fmt.Sprintf("%s %s", scheme, credential))
	return r
}

func (r *Request) SetBasicAuth(username, password string) *Request {
	auth := username + ":" + password
	credential := base64.StdEncoding.EncodeToString([]byte(auth))
	r.SetAuth("Basic", credential)
	return r
}

func (r *Request) SetBearerToken(token string) *Request {
	return r.SetAuth("Bearer", token)
}
