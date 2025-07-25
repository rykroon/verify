package httpx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Response struct {
	*http.Response
	BodyBytes []byte
}

func (r *Response) IsInformational() bool {
	return r.StatusCode >= 100 && r.StatusCode < 200
}

func (r *Response) IsSuccess() bool {
	return r.StatusCode >= 200 && r.StatusCode < 300
}

func (r *Response) IsRedirect() bool {
	return r.StatusCode >= 300 && r.StatusCode < 400
}

func (r *Response) IsClientError() bool {
	return r.StatusCode >= 400 && r.StatusCode < 500
}

func (r *Response) IsServerError() bool {
	return r.StatusCode >= 500
}

func (r *Response) IsError() bool {
	return r.StatusCode >= 400
}

func (r *Response) ContentType() string {
	return r.Header.Get("Content-Type")
}

func (r *Response) ToString() string {
	return string(r.BodyBytes)
}

func (r *Response) ToForm() (url.Values, error) {
	if !strings.HasPrefix(r.ContentType(), "application/x-www-form-urlencoded") {
		return nil, fmt.Errorf("unexpected content-type: %s", r.ContentType())
	}
	return url.ParseQuery(string(r.BodyBytes))
}

func (r *Response) ToJson(v any) error {
	if !strings.HasPrefix(r.ContentType(), "application/json") {
		return fmt.Errorf("unexpected content-type: %s", r.ContentType())
	}
	return json.Unmarshal(r.BodyBytes, v)
}
