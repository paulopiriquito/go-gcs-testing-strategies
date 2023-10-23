package main

import (
	"context"
	application "gcs-testing-strategies/app"
	"gcs-testing-strategies/gcs"
	"os"
)

func main() {
	ctx := context.Background()
	testing := os.Getenv("ENV") == "TEST"
	fileName := "test-file"

	bucketClient := loadClient(ctx, testing)

	if testing {
		application.Test(fileName, bucketClient, ctx)
	} else {
		application.App(fileName, bucketClient, ctx)
	}
}

func loadClient(ctx context.Context, testing bool) application.BucketClient {
	gcsClient, err := gcs.NewStorageClient(ctx, testing)
	if err != nil {
		panic(err)
	}
	bucketClient := &gcs.BucketClient{
		BucketName:          os.Getenv("GCS_BUCKET_NAME"),
		GoogleStorageClient: gcsClient,
	}
	return bucketClient
}
