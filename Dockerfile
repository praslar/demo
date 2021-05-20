FROM golang:alpine as build-env

ARG VERSION=0.0.0
ARG SERVICE="svc-general"

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# cache dependencies first
WORKDIR /svc
COPY go.mod /svc
COPY go.sum /svc
RUN go mod download

# lastly copy source, any change in source will not break above cache
COPY . /svc

# Build the binary
RUN go build -a -ldflags="-s -w -X 'main.version=${VERSION}' -X 'main.name=${SERVICE}'" -o /app ./main.go

# # <- Second step to build minimal image
FROM alpine:3.11

# RUN apk add --no-cache git ca-certificates tzdata

# we have no self-sign certificate, don't need to update
# && update-ca-certificates
WORKDIR /svc
COPY ./conf/app.conf /svc/conf/app.conf
COPY --from=build-env /app /svc/app

ENTRYPOINT ["/svc/app"]
