package sinch

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"

	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/rykroon/verify/internal/utils"
)

type client struct {
	applicationKey    string
	applicationSecret string
}

func NewClient(applicationKey, applicationSecret string) *client {
	return &client{applicationKey, applicationSecret}
}

func (c *client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	u, err := url.JoinPath("https://verification.api.sinch.com/verification/v1", path)
	if err != nil {
		return nil, fmt.Errorf("failed to join path: %w", err)
	}

	var copy1 io.Reader
	var copy2 io.Reader

	if body != nil {
		var buf bytes.Buffer
		_, err = io.Copy(&buf, body)
		if err != nil {
			return nil, err
		}

		copy1 = bytes.NewReader(buf.Bytes())
		copy2 = bytes.NewReader(buf.Bytes())
	}

	req, err := http.NewRequest(method, u, copy1)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	c.signRequest(req, copy2)
	return req, nil

}

func (c *client) handleResponse(resp *utils.CachedResponse) (json.RawMessage, error) {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("http error: code: %d, body: %s", resp.StatusCode, string(resp.Body))
	}

	var result json.RawMessage
	if err := json.Unmarshal(resp.Body, &result); err != nil {
		return nil, fmt.Errorf("failed to decode resposne body as json: %w", err)
	}

	return result, nil
}

func (c *client) signRequest(req *http.Request, body io.Reader) error {
	contentMD5 := ""
	contentType := req.Header.Get("Content-Type")
	if body != nil {
		data, err := io.ReadAll(body)
		if err != nil {
			return fmt.Errorf("failed to read body: %w", err)
		}
		sum := md5.Sum(data)
		contentMD5 = base64.StdEncoding.EncodeToString(sum[:])
	}

	t := time.Now().UTC()
	timestamp := t.Format("2006-01-02T15:04:05.0000000Z")
	req.Header.Add("x-timestamp", timestamp)

	stringToSign := req.Method + "\n" +
		contentMD5 + "\n" +
		contentType + "\n" +
		fmt.Sprintf("x-timestamp:%s", timestamp) + "\n" +
		req.URL.Path

	key, err := base64.StdEncoding.DecodeString(c.applicationSecret)
	if err != nil {
		return fmt.Errorf("failed to decode string: %w", err)
	}

	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(stringToSign))
	digest := mac.Sum(nil)
	sig := base64.StdEncoding.EncodeToString(digest)

	req.Header.Set("Authorization", fmt.Sprintf("Application %s:%s", c.applicationKey, sig))
	return nil
}
