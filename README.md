# go-gcs-testing-strategies
A GO repo with GCS (Google Cloud Storage) unit and integration tests strategies

This project is a demonstration of how to emulate GCS (Google Cloud Storage) in GO unit testing and how to perform
integration tests in a docker stack.

## Requirements
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Golang](https://golang.org/doc/install) (optional)

## Usage

### Integration testing

To run the integration tests, you need to have docker and docker-compose installed. Then, you can run the following int the repo root:
```bash
docker-compose -f ./integration-tests/docker-compose.yml up --build --exit-code-from gcs-client-tester
```
This will build and run this GO appplication in "TEST" environment, which will use the GCS emulator to perform the tests.
The test consists of:
- Uploading a file to the bucket
- Listing the files in the bucket
- Downloading the file from the bucket
- Deleting the file from the bucket

Once the tests are finished, the program outputs "Test passed" and exits with code 0.

To remove the docker stack, run:
```bash
docker-compose -f ./integration-tests/docker-compose.yml down --remove-orphans --volumes
```

### Unit testing

To run the unit tests, you need to have Golang installed. Then, you can run the following int the repo root:
```bash
cd ./src
go test ./... -v
```

The unit test use a in-memory implementation of the GCS client, which is not suitable for integration testing.
However it is useful for unit testing and for debugging purposes.
The in-memory implementation is compatible with the GCS client interface, so it can be used as a drop-in replacement
for the real GCS client.


## Extending the application
The application defines a simple interface for the GCS client, which can be implemented by any GCS client.
You can implement the interface with the real GCS client, or with a mock client, or with an in-memory client.
The interface is located in `./src/app/bucket_client.go` and it is called `BucketClient`.
The present implementation of the interface is located in `./src/gcs/gcs_client.go`.

It currently supports the following operations:
- WriteFile
- ListObjects
- ReadFile
- DeleteFile

The implementation also has a constructor function `NewGCSClient` which returns a new instance of the GCS client with the
integration test configuration support.

### Environment variables

The implementation dependency on the environment variable `GOOGLE_APPLICATION_CREDENTIALS` and/or
`STORAGE_EMULATOR_HOST`. If you want to run the integration tests in a different environment,
without emulation, you need to set `GOOGLE_APPLICATION_CREDENTIALS` to a valid GCP service account key file location
and `STORAGE_EMULATOR_HOST` must not be present.

If you only want to test the application in a simulated environment, you must not set `GOOGLE_APPLICATION_CREDENTIALS`
and set `STORAGE_EMULATOR_HOST` to the emulator host.

The second scenario is the default one for the purpose of this demonstration.