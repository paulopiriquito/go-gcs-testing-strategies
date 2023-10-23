## Builder Layer
FROM golang:1.21-bookworm AS build

WORKDIR /project

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go build -o /app

## Deployed Layer
FROM alpine AS run

WORKDIR /

RUN touch .env
RUN mkdir /output

COPY --from=build /etc/passwd /etc/passwd
COPY --from=build /app /app

## Golang runtime dependency
RUN apk add --no-cache gcompat

## Add executable permissions to app file owners group
RUN chmod -R 777 /app
## Make application-runners group owner of app binaries file
RUN chown -R root:root /app

ENTRYPOINT ["/app"]