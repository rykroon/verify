package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
)

type Content struct {
	Data    []byte
	Type    string
	Charset map[string]string
}

func (c Content) IsJson() bool {
	return c.Type == "application/json"
}

func (c *Content) DecodeJsonInto(v any) error {
	return json.NewDecoder(bytes.NewReader(c.Data)).Decode(v)
}

func SendRequest(client *http.Client, req *http.Request) (*Content, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http error: status code: %d, response body: %s", resp.StatusCode, string(data))
	}

	mediatype, params, err := mime.ParseMediaType(resp.Header.Get("Content-Type"))
	if err != nil {
		return nil, fmt.Errorf("failed to parse media type: %w", err)
	}

	return &Content{Data: data, Type: mediatype, Charset: params}, nil
}
