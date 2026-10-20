FROM golang:1.26.10-alpine3.24 AS build_base

RUN apk add --no-cache git
RUN apk add build-base

# Set the Current Working Directory inside the container
WORKDIR /tmp/app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o ./main .

FROM alpine:3.24
RUN apk add --no-cache ca-certificates

COPY --from=build_base /tmp/app/main /main

# Run the startup script
CMD ["/main"]
