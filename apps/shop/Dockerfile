#build stage
FROM golang:1.20 AS builder
ARG SERVICE_NAME 
WORKDIR /go/src/sds
ADD . .
WORKDIR /go/src/sds/apps
RUN go get sds-final
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app ${SERVICE_NAME}/main.go

#final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/sds/apps .
CMD ["./app"]