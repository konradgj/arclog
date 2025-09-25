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
	ReasonNone      UploadReason = ""
	ReasonCreate    UploadReason = "create"
	ReasonQueueFull UploadReason = "queue_full"
)
