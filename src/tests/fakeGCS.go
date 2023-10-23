package tests

import (
	"cloud.google.com/go/storage"
	"context"
	"gcs-testing-strategies/app"
	"gcs-testing-strategies/gcs"
	"github.com/fsouza/fake-gcs-server/fakestorage"
	"os"
)

func GetBucketClient(bucketName string, context context.Context) (app.BucketClient, *fakestorage.Server) {
	bucketClient, server := getFakeGCSClient(bucketName, context)

	client := gcs.BucketClient{
		GoogleStorageClient: bucketClient,
		BucketName:          bucketName,
	}

	return client, server
}

func getFakeGCSClient(bucketName string, context context.Context) (*storage.Client, *fakestorage.Server) {
	server, err := fakestorage.NewServerWithOptions(fakestorage.Options{Scheme: "http"})
	if err != nil {
		panic(err)
	}
	server.CreateBucketWithOpts(fakestorage.CreateBucketOpts{Name: bucketName})

	err = os.Setenv("STORAGE_EMULATOR_HOST", server.URL())
	if err != nil {
		return nil, nil
	}

	client, err := gcs.NewStorageClient(context, true)
	if err != nil {
		panic(err)
	}
	return client, server
}
