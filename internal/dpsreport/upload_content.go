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

	"github.com/konradgj/arclog/internal/db"
)

const uploadContentEndpoint = "uploadContent"

type UploadContentOptions struct {
	UserToken   string
	Anonymous   bool
	DetailedWvW bool
}

func (c *Client) UploadContent(filePath string, opts UploadContentOptions) (*UploadResponse, error) {
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%w: %v", db.ErrFileMissing, err)
		}
		return nil, fmt.Errorf("%w: %v", db.ErrInternal, err)
	}
	defer file.Close()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", db.ErrInternal, err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return nil, fmt.Errorf("%w: %v", db.ErrInternal, err)
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
		return nil, fmt.Errorf("%w: %v", db.ErrInternal, err)
	}

	url := baseUrl + uploadContentEndpoint
	req, err := http.NewRequest("POST", url, &body)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", db.ErrInternal, err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", db.ErrHttp, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		b, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("%w: unexpected status %d: %s", db.ErrHttp, resp.StatusCode, string(b))
	}

	var uploadResponse UploadResponse
	if err := json.NewDecoder(resp.Body).Decode(&uploadResponse); err != nil {
		return &uploadResponse, fmt.Errorf("%w: %v", db.ErrDecode, err)
	}

	if uploadResponse.Error != nil {
		return &uploadResponse, fmt.Errorf("%w: %s", db.ErrServerError, *uploadResponse.Error)
	}

	return &uploadResponse, nil
}
