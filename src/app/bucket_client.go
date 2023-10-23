package app

import "context"

type BucketClient interface {
	WriteFile(file File, ctx context.Context) error
	ReadFile(fileName string, ctx context.Context) (File, error)
	DeleteFile(filename string, ctx context.Context) error
	ListObjects(ctx context.Context) ([]string, error)
}

type File struct {
	Name    string
	Content []byte
}
