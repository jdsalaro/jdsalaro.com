
# STAGE: BUILDER
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN apk add --no-cache git
RUN go get ./...
RUN go build -o main main.go

# STAGE: FINAL
FROM alpine:latest
RUN apk add --no-cache bash
COPY --from=builder /go/src/app/main main

LABEL Name=app
EXPOSE 8080

RUN ["chmod", "+x", "main"]
ENTRYPOINT ["./main"]