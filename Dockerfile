# Build the Go binary
FROM golang:1.25-alpine AS goapp
WORKDIR /app

RUN apk --no-cache add curl make

COPY go.mod go.sum ./
RUN go mod download

COPY Makefile .
COPY main.go  .
COPY internal/ internal/

RUN make
RUN go build -o ./goapp

# Build the final image
FROM alpine:latest AS release
WORKDIR /app

COPY --from=goapp /app/goapp /goapp
# database data at /pb/pb_data - mount as volume

EXPOSE 8080
ENV PUBLIC_DIR=/dist
CMD ["/goapp", "serve", "--http=0.0.0.0:8080"]