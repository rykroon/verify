package httpx

import (
	"io"
	"net/http"
)

type ResponseBodyProvider interface {
	ContentType() string
	WriteTo(io.Writer) (int64, error)
}

func WriteBody(w http.ResponseWriter, p ResponseBodyProvider) error {
	w.Header().Set("Content-Type", p.ContentType())
	_, err := p.WriteTo(w)
	return err
}

func IsInformational(r *http.Response) bool { return r.StatusCode >= 100 && r.StatusCode < 200 }
func IsSuccess(r *http.Response) bool       { return r.StatusCode >= 200 && r.StatusCode < 300 }
func IsRedirect(r *http.Response) bool      { return r.StatusCode >= 300 && r.StatusCode < 400 }
func IsClientError(r *http.Response) bool   { return r.StatusCode >= 400 && r.StatusCode < 500 }
func IsServerError(r *http.Response) bool   { return r.StatusCode >= 500 }
func IsError(r *http.Response) bool         { return r.StatusCode >= 400 }

func ReadBodyFromResponse(resp *http.Response) (*Body, error) {
	return ReadBody(resp.Header.Get("Content-Type"), resp.Body)
}
