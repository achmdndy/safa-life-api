package storage

import (
	"context"
	"fmt"
	"io"
	"net/url"
	"strings"

	domStorage "github.com/safalife/core-api/src/domain/storage"

	"github.com/aws/aws-sdk-go-v2/aws"
	awscfg "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// S3StorageService implements StorageServiceInterface using AWS S3 (or S3-compatible) backend
type S3StorageService struct {
	client        *s3.Client
	bucket        string
	region        string
	endpoint      string
	usePathStyle  bool
	publicURLBase string
}

// NewS3StorageService constructs the S3 storage service
func NewS3StorageService(ctx context.Context, cfg S3Config) (*S3StorageService, error) {
	if !cfg.Enabled {
		return nil, fmt.Errorf("s3 storage is disabled in configuration")
	}
	if cfg.Bucket == "" || cfg.Region == "" {
		return nil, fmt.Errorf("s3 configuration requires non-empty bucket and region")
	}

	awsCfg, err := awscfg.LoadDefaultConfig(
		ctx,
		awscfg.WithRegion(cfg.Region),
		awscfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.AccessKeyID, cfg.SecretAccessKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.UsePathStyle = cfg.UsePathStyle
		if cfg.Endpoint != "" {
			o.BaseEndpoint = aws.String(cfg.Endpoint)
		}
	})

	return &S3StorageService{
		client:        client,
		bucket:        cfg.Bucket,
		region:        cfg.Region,
		endpoint:      cfg.Endpoint,
		usePathStyle:  cfg.UsePathStyle,
		publicURLBase: strings.TrimRight(cfg.PublicURLBase, "/"),
	}, nil
}

// Upload stores an object, returning its metadata
func (s *S3StorageService) Upload(ctx context.Context, in domStorage.UploadInput) (domStorage.ObjectInfo, error) {
	if in.Key == "" || in.Body == nil {
		return domStorage.ObjectInfo{}, fmt.Errorf("upload requires non-empty key and body")
	}

	put := &s3.PutObjectInput{
		Bucket:   aws.String(s.bucket),
		Key:      aws.String(in.Key),
		Body:     in.Body,
		Metadata: in.Metadata,
	}
	if in.ContentType != "" {
		put.ContentType = aws.String(in.ContentType)
	}

	resp, err := s.client.PutObject(ctx, put)
	if err != nil {
		return domStorage.ObjectInfo{}, err
	}

	return domStorage.ObjectInfo{
		Key:  in.Key,
		URL:  s.GetURL(in.Key),
		ETag: aws.ToString(resp.ETag),
	}, nil
}

// Get retrieves an object, returning its body and basic metadata
func (s *S3StorageService) Get(ctx context.Context, key string) (io.ReadCloser, *domStorage.ObjectInfo, error) {
	if key == "" {
		return nil, nil, fmt.Errorf("get requires non-empty key")
	}
	out, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, nil, err
	}
	info := &domStorage.ObjectInfo{
		Key:         key,
		URL:         s.GetURL(key),
		ETag:        aws.ToString(out.ETag),
		Size:        aws.ToInt64(out.ContentLength),
		ContentType: aws.ToString(out.ContentType),
	}
	if out.LastModified != nil {
		info.LastModified = *out.LastModified
	}
	return out.Body, info, nil
}

// Delete removes an object
func (s *S3StorageService) Delete(ctx context.Context, key string) error {
	if key == "" {
		return fmt.Errorf("delete requires non-empty key")
	}
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})
	return err
}

// List lists objects with optional pagination
func (s *S3StorageService) List(ctx context.Context, in domStorage.ListInput) ([]domStorage.ObjectInfo, string, error) {
	maxKeys := int32(100)
	if in.MaxKeys > 0 {
		maxKeys = in.MaxKeys
	}
	out, err := s.client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket:  aws.String(s.bucket),
		Prefix:  aws.String(in.Prefix),
		MaxKeys: aws.Int32(maxKeys),
		ContinuationToken: func() *string {
			if in.ContinuationToken != "" {
				return aws.String(in.ContinuationToken)
			}
			return nil
		}(),
	})
	if err != nil {
		return nil, "", err
	}
	items := make([]domStorage.ObjectInfo, 0, len(out.Contents))
	for _, o := range out.Contents {
		key := aws.ToString(o.Key)
		items = append(items, domStorage.ObjectInfo{
			Key:          key,
			URL:          s.GetURL(key),
			ETag:         aws.ToString(o.ETag),
			Size:         aws.ToInt64(o.Size),
			LastModified: aws.ToTime(o.LastModified),
		})
	}
	next := aws.ToString(out.NextContinuationToken)
	return items, next, nil
}

// GetURL returns a publicly accessible URL for the given object key
func (s *S3StorageService) GetURL(key string) string {
	key = strings.TrimLeft(key, "/")
	if s.publicURLBase != "" {
		return s.publicURLBase + "/" + url.PathEscape(key)
	}

	if s.endpoint != "" {
		base := strings.TrimRight(s.endpoint, "/")
		if s.usePathStyle {
			return fmt.Sprintf("%s/%s/%s", base, s.bucket, url.PathEscape(key))
		}
		return fmt.Sprintf("%s/%s/%s", base, s.bucket, url.PathEscape(key))
	}

	if s.usePathStyle {
		return fmt.Sprintf("https://s3.%s.amazonaws.com/%s/%s", s.region, s.bucket, url.PathEscape(key))
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, url.PathEscape(key))
}

// Ensure S3StorageService implements the interface
var _ domStorage.StorageServiceInterface = (*S3StorageService)(nil)
