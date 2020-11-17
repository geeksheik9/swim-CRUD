FROM golang:alpine AS builder
ARG VERSION
WORKDIR /go/src/swim-CRUD
COPY . .
WORKDIR /go/src/swim-CRUD/cmd/svr
RUN go mod download
RUN go get -d -v; \
    go install -v; \
    go build -gcflags "all=-N -l" -ldflags "-X main.version=${VERSION}" -o app;

FROM alpine:latest
WORKDIR /root
COPY --from=builder /go/src/swim-CRUD/cmd/svr/app .

COPY ./swagger-ui /root/swagger-ui

CMD ["./app"]
