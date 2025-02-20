package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ContentType string

const (
	ContentTypeJSON ContentType = "application/json"
	ContentTypeForm ContentType = "application/x-www-form-urlencoded"
)

type Request struct {
	httpClient *http.Client
}

type DefaultApiResponse struct {
	Ok     bool            `json:"ok"`
	Result json.RawMessage `json:"result"`
}

func processRequest(resp *http.Response, outModel interface{}) error {
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error response: %s, body: %s", resp.Status, string(respBody))
	}
	var apiResponse DefaultApiResponse
	err = json.Unmarshal(respBody, &apiResponse)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response body: %v", err)
	}

	return json.Unmarshal(apiResponse.Result, outModel)

}

func (r *Request) Post(url string, contentType ContentType, body interface{}, outModel interface{}) error {
	var bodyReader io.Reader
	if body != nil && !isEmpty(body) {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}
	resp, err := http.Post(url, string(contentType), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %v", err)
	}
	defer resp.Body.Close()
	return processRequest(resp, outModel)
}
