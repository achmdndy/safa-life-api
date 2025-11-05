package storage

import (
	"context"
	"io"
	"time"
)

// ObjectInfo represents metadata of an object stored in S3-compatible storage
type ObjectInfo struct {
	Key          string
	URL          string
	ETag         string
	Size         int64
	ContentType  string
	LastModified time.Time
}

// UploadInput represents parameters for uploading an object
type UploadInput struct {
	Key         string
	Body        io.Reader
	ContentType string
	PublicRead  bool
	Metadata    map[string]string
}

// ListInput represents parameters for listing objects by prefix
type ListInput struct {
	Prefix            string
	MaxKeys           int32
	ContinuationToken string
}

// StorageServiceInterface defines CRUD operations for object storage
type StorageServiceInterface interface {
	Upload(ctx context.Context, in UploadInput) (ObjectInfo, error)
	Get(ctx context.Context, key string) (io.ReadCloser, *ObjectInfo, error)
	Delete(ctx context.Context, key string) error
	List(ctx context.Context, in ListInput) ([]ObjectInfo, string, error)
	GetURL(key string) string
}
