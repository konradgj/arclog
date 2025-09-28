package db

import "errors"

type UploadStatus string

const (
	StatusPending   UploadStatus = "pending"
	StatusUploading UploadStatus = "uploading"
	StatusUploaded  UploadStatus = "uploaded"
	StatusFailed    UploadStatus = "failed"
	StatusSkipped   UploadStatus = "skipped"
)

type UploadReason string

const (
	ReasonUnknown     UploadReason = "unkown"
	ReasonCreate      UploadReason = "create"
	ReasonUploading   UploadReason = "uploading"
	ReasonSuccess     UploadReason = "success"
	ReasonFileMissing UploadReason = "file_missing"
	ReasonInternal    UploadReason = "error_internal"
	ReasonHttp        UploadReason = "error_http"
	ReasonDecode      UploadReason = "error_internal"
	ReasonServerError UploadReason = "error_server"
)

var (
	ErrFileMissing = errors.New("file missing")
	ErrInternal    = errors.New("internal error")
	ErrHttp        = errors.New("http error")
	ErrDecode      = errors.New("decode error")
	ErrServerError = errors.New("server error")
)

func ErrMapToReason(err error) string {
	var reason string

	switch {
	case errors.Is(err, ErrFileMissing):
		reason = string(ReasonFileMissing)
	case errors.Is(err, ErrInternal):
		reason = string(ReasonInternal)
	case errors.Is(err, ErrHttp):
		reason = string(ReasonHttp)
	case errors.Is(err, ErrDecode):
		reason = string(ReasonDecode)
	case errors.Is(err, ErrServerError):
		reason = string(ReasonServerError)
	default:
		reason = string(ReasonUnknown)
	}
	return reason
}
