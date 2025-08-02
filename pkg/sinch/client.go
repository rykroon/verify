package sinch

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"

	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/rykroon/verify/internal/httpx"
)

type client struct {
	applicationKey    string
	applicationSecret string
	httpClient        *http.Client
}

func NewClient(applicationKey, applicationSecret string) *client {
	return &client{applicationKey, applicationSecret, http.DefaultClient}
}

func (c *client) SetHttpClient(client *http.Client) {
	c.httpClient = client
}

func (c *client) newRequest(method, path string, body httpx.RequestBodyProvider) (*http.Request, error) {
	u, err := url.JoinPath("https://verification.api.sinch.com/verification/v1", path)
	if err != nil {
		return nil, fmt.Errorf("failed to join path: %w", err)
	}

	req, err := httpx.NewRequest(method, u, body)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.signRequest(req, body)
	return req, nil

}

func (c *client) sendRequest(req *http.Request) (*httpx.Body, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	respBody, err := httpx.ReadBodyFromResponse(resp)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if httpx.IsError(resp) {
		return nil, fmt.Errorf("http error: code: %d, body: %s", resp.StatusCode, respBody.ToString())
	}

	return respBody, nil
}

func (c *client) signRequest(req *http.Request, body httpx.RequestBodyProvider) error {
	contentMD5 := ""
	contentType := ""
	if body != nil {
		data, err := io.ReadAll(body.Reader())
		if err != nil {
			return fmt.Errorf("failed to read body: %w", err)
		}
		sum := md5.Sum(data)
		contentMD5 = base64.StdEncoding.EncodeToString(sum[:])
		contentType = body.ContentType()
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
