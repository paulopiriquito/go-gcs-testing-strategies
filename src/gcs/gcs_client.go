package gcs

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"gcs-testing-strategies/app"
	"google.golang.org/api/option"
	"os"
)

type BucketClient struct {
	BucketName          string
	GoogleStorageClient *storage.Client
}

func NewStorageClient(ctx context.Context, testing bool) (*storage.Client, error) {
	if testing {
		endpoint := os.Getenv("STORAGE_EMULATOR_HOST")
		if endpoint == "" {
			return nil, fmt.Errorf("STORAGE_EMULATOR_HOST not set")
		}
		return storage.NewClient(ctx, option.WithoutAuthentication(), storage.WithJSONReads())
	} else {
		return storage.NewClient(ctx)
	}
}

func (bucketClient BucketClient) ListObjects(ctx context.Context) ([]string, error) {
	objects := bucketClient.GoogleStorageClient.
		Bucket(bucketClient.BucketName).
		Objects(ctx, nil)

	var objectNames []string
	for {
		attrs, err := objects.Next()
		if err != nil {
			break
		}
		objectNames = append(objectNames, attrs.Name)
	}
	return objectNames, nil
}

func (bucketClient BucketClient) DeleteFile(filename string, ctx context.Context) error {
	return bucketClient.GoogleStorageClient.
		Bucket(bucketClient.BucketName).
		Object(filename).
		Delete(ctx)
}

func (bucketClient BucketClient) WriteFile(file app.File, ctx context.Context) error {
	writer := bucketClient.GoogleStorageClient.Bucket(bucketClient.BucketName).Object(file.Name).NewWriter(ctx)
	written, err := writer.Write(file.Content)
	defer func(writer *storage.Writer) {
		err := writer.Close()
		if err != nil {
			panic(err)
		}
	}(writer)

	if err != nil {
		return err
	}
	if written != len(file.Content) {
		return fmt.Errorf("could not write file")
	}
	return nil
}

func (bucketClient BucketClient) ReadFile(fileName string, ctx context.Context) (app.File, error) {
	reader, err := bucketClient.GoogleStorageClient.
		Bucket(bucketClient.BucketName).
		Object(fileName).
		NewReader(ctx)

	if err != nil {
		return app.File{}, err
	}

	bucketFile := app.File{
		Name: fileName,
	}
	bucketFile.Content = make([]byte, reader.Attrs.Size)
	bytesRead, err := reader.Read(bucketFile.Content)
	if err != nil && bytesRead != len(bucketFile.Content) {
		return app.File{}, err
	}
	err = reader.Close()
	if err != nil {
		return app.File{}, err
	}

	return bucketFile, nil
}
