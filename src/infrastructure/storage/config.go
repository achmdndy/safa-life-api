package storage

// S3Config holds configuration for connecting to AWS S3 or S3-compatible storage
type S3Config struct {
	Enabled         bool
	Bucket          string
	Region          string
	AccessKeyID     string
	SecretAccessKey string
	Endpoint        string
	UsePathStyle    bool
	PublicURLBase   string
}
