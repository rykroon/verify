package httpx

import (
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	*http.Response
	body []byte
	err  error
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

func (r *Response) mediaType() string {
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

func (r *Response) IsPlainText() bool {
	return r.mediaType() == "text/plain"
}

func (r *Response) IsJson() bool {
	return r.ContentType() == "application/json"
}

func (r *Response) IsForm() bool {
	return r.ContentType() == "application/x-www-form-urlencoded"
}

func (r *Response) ReadBody() ([]byte, error) {
	if r.body != nil || r.err != nil {
		return r.body, r.err
	}
	defer r.Body.Close()
	r.body, r.err = io.ReadAll(r.Body)
	return r.body, r.err
}

func (r *Response) ToString() (string, error) {
	body, err := r.ReadBody()
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (r *Response) ToForm() (url.Values, error) {
	body, err := r.ReadBody()
	if err != nil {
		return nil, err
	}
	return url.ParseQuery(string(body))
}

func (r *Response) ToJson(v any) error {
	body, err := r.ReadBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
