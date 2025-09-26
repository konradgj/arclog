package db

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
	ReasonNone      UploadReason = "none"
	ReasonCreate    UploadReason = "create"
	ReasonUploading UploadReason = "uploading"
	ReasonSuccess   UploadReason = "success"
	ReasonQueueFull UploadReason = "queue_full"
)
