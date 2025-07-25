package telnyx

import (
	"io"
	"net/http"
	"net/url"

	"github.com/rykroon/verify/internal/httpx"
)

type Client struct {
	apiToken string
}

func NewClient(apiToken string) *Client {
	return &Client{apiToken: apiToken}
}

func (c *Client) sendRequest(method, path string, builder httpx.BodyBuilder) (*httpx.Response, error) {
	urlStr, err := url.JoinPath("https://api.telnyx.com/v2", path)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if builder != nil {
		reader, err := builder.ToReader()
		if err != nil {
			return nil, err
		}
		body = reader
	}

	req, err := httpx.NewRequest(method, urlStr, body)

	if builder != nil {
		req.SetContentType(builder.ContentType())
	}

	req.SetBearerToken(c.apiToken)

	resp, err := http.DefaultClient.Do(req.Request)
	if err != nil {
		return nil, err
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &httpx.Response{Response: resp, BodyBytes: bodyBytes}, nil
}

// type Record interface {
// 	GetRecordType() string
// }

// type ApiResponse[T Record] struct {
// 	Data T `json:"data"`
// }
