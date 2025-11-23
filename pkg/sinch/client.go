package sinch

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"

	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient        *http.Client
	applicationKey    string
	applicationSecret string
}

func NewClient(httpClient *http.Client, applicationKey, applicationSecret string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &Client{httpClient, applicationKey, applicationSecret}
}

func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	u, err := url.JoinPath("https://verification.api.sinch.com/verification/v1", path)
	if err != nil {
		return nil, fmt.Errorf("failed to join path: %w", err)
	}

	// copy the body so that it can be used for signing and sending the request.
	var buf bytes.Buffer
	if body != nil {
		_, err = io.Copy(&buf, body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u, bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	c.signRequest(req, buf.Bytes())
	return req, nil

}

func (c *Client) signRequest(req *http.Request, body []byte) error {
	contentMD5 := ""
	contentType := req.Header.Get("Content-Type")
	if body != nil {
		sum := md5.Sum(body)
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
