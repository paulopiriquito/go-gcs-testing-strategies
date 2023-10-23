package app

import "context"

func Test(fileName string, bucket BucketClient, ctx context.Context) {
	App(fileName, bucket, ctx)
	err := bucket.DeleteFile(fileName, ctx)
	if err != nil {
		panic(err)
	}
	println("Test passed")
}

func App(fileName string, bucket BucketClient, ctx context.Context) {
	file := File{
		Name:    fileName,
		Content: []byte("test-content"),
	}
	err := bucket.WriteFile(file, ctx)
	if err != nil {
		panic(err)
	}
	files, err := bucket.ListObjects(ctx)
	if len(files) != 1 || err != nil {
		panic("File not written")
	}
	file, err = bucket.ReadFile(fileName, ctx)
	if err != nil {
		panic(err)
	}
}
