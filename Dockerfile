
#build stage
FROM golang:1.12.5 AS builder
WORKDIR /go/src/app
COPY . .
# RUN apk add --no-cache git
RUN go get -d -v ./...
RUN go install -v ./...

#final stage
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates
# COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=children_app_api Version=0.0.1
EXPOSE 8000
