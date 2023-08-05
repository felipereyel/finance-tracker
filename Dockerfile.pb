FROM golang:1.20-alpine AS goapp
WORKDIR /app

RUN go install github.com/cosmtrek/air@latest
COPY .air.toml ./

COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml", "serve"]