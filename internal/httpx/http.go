package httpx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpError struct {
	StatusCode  int
	Body        []byte
	ContentType string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("http error: %d", e.StatusCode)
}

func HandleResponse(resp *http.Response, result any) error {
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read body: %w", err)
	}

	contentType := resp.Header.Get("Content-Type")

	if resp.StatusCode >= 400 {
		return &HttpError{
			StatusCode:  resp.StatusCode,
			Body:        body,
			ContentType: contentType,
		}
	}

	if result != nil {
		if err := json.Unmarshal(body, result); err != nil {
			return fmt.Errorf("json unmarshal: %w", err)
		}
	}

	return nil
}
