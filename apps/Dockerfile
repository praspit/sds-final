#build stage
FROM golang:1.20 AS builder
WORKDIR /go/src/sds
ADD . .
RUN go get sds-final
RUN echo $(ls)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-extldflags "-static"' -o app apps/shop/main.go

#final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/sds/app .
CMD ["./app"]