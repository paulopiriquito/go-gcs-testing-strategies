package app

import (
	"context"
	bucket "gcs-testing-strategies/app"
	fakeServer "gcs-testing-strategies/tests"
	"testing"
)

func Test_WriteToBucket(t *testing.T) {
	client, server := fakeServer.GetBucketClient("test-bucket", context.Background())
	t.Cleanup(server.Stop)

	fileName := "test-write-file"
	fileContent := "test-content"

	bucketFileToWrite := bucket.File{
		Name:    fileName,
		Content: []byte(fileContent),
	}

	err := client.WriteFile(bucketFileToWrite, context.Background())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	files, err := client.ListObjects(context.Background())
	if len(files) != 1 || err != nil {
		t.Error("File not written")
		t.Fail()
	}
}

func Test_ReadFromBucket(t *testing.T) {
	client, server := fakeServer.GetBucketClient("test-bucket-read", context.Background())
	t.Cleanup(server.Stop)

	fileName := "test-read-file"
	expectedFileContent := "test-content"

	bucketFileToWrite := bucket.File{
		Name:    fileName,
		Content: []byte(expectedFileContent),
	}

	err := client.WriteFile(bucketFileToWrite, context.Background())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	bucketFile, err := client.ReadFile(fileName, context.Background())
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if bucketFile.Name != fileName {
		t.Error("File name is not correct")
		t.Fail()
	}
	if string(bucketFile.Content) != expectedFileContent {
		t.Error("File content is not correct")
		t.Fail()
	}
}

func Test_DeleteFromBucket(t *testing.T) {
	client, server := fakeServer.GetBucketClient("test-bucket-read", context.Background())
	t.Cleanup(server.Stop)

	fileName := "test-read-file"
	expectedFileContent := "test-content"

	bucketFileToWrite := bucket.File{
		Name:    fileName,
		Content: []byte(expectedFileContent),
	}

	err := client.WriteFile(bucketFileToWrite, context.Background())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	err = client.DeleteFile(fileName, context.Background())
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	files, err := client.ListObjects(context.Background())
	if len(files) != 0 || err != nil {
		t.Error("File not deleted")
		t.Fail()
	}
}
