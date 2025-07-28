package httpx

import (
	"io"
	"mime"
	"net/http"
	"strings"
)

type Response struct {
	*http.Response
	body        *Body
	readBodyErr error
}

func WrapResponse(resp *http.Response) *Response {
	return &Response{Response: resp}
}

func (r *Response) IsInformational() bool { return r.StatusCode >= 100 && r.StatusCode < 200 }
func (r *Response) IsSuccess() bool       { return r.StatusCode >= 200 && r.StatusCode < 300 }
func (r *Response) IsRedirect() bool      { return r.StatusCode >= 300 && r.StatusCode < 400 }
func (r *Response) IsClientError() bool   { return r.StatusCode >= 400 && r.StatusCode < 500 }
func (r *Response) IsServerError() bool   { return r.StatusCode >= 500 }
func (r *Response) IsError() bool         { return r.StatusCode >= 400 }

func (r *Response) ContentType() string {
	return r.Header.Get("Content-Type")
}

func (r *Response) MediaType() string {
	ctype := r.ContentType()
	if ctype == "" {
		return ""
	}
	mt, _, err := mime.ParseMediaType(ctype)
	if err != nil {
		return strings.ToLower(ctype) // fallback
	}
	return strings.ToLower(mt)
}

func (r *Response) IsJson() bool {
	return r.MediaType() == "application/json"
}

func (r *Response) ReadBody() (*Body, error) {
	if r.body != nil || r.readBodyErr != nil {
		return r.body, r.readBodyErr
	}
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		r.readBodyErr = err
		return nil, r.readBodyErr
	}
	r.body = NewBody(data, r.ContentType())
	return r.body, err
}
