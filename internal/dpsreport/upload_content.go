package dpsreport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

const uploadContentEndpoint = "uploadContent"

type UploadContentOptions struct {
	UserToken   string
	Anonymous   bool
	DetailedWvW bool
}

func (c *Client) uploadContent(filePath string, opts UploadContentOptions) (*UploadResponse, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("could not create form file: %w", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("could not copy file: %w", err)
	}

	_ = writer.WriteField("json", "1")
	_ = writer.WriteField("generator", "ei")
	if opts.UserToken != "" {
		_ = writer.WriteField("userToken", opts.UserToken)
	}
	if opts.Anonymous {
		_ = writer.WriteField("anonymous", "true")
	}
	if opts.DetailedWvW {
		_ = writer.WriteField("detailedwvw", "true")
	}

	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("could not close writer: %w", err)
	}

	url := baseUrl + uploadContentEndpoint
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("upload request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("unexpected status %d: %s", resp.StatusCode, string(b))
	}

	var uploadResponse UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return &uploadResponse, fmt.Errorf("could not decode response: %w", err)
	}

	if uploadResponse.Error != nil {
		return &uploadResponse, fmt.Errorf("server error: %s", *uploadResponse.Error)
	}

	return &uploadResponse, nil
}
