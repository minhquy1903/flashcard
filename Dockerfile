FROM golang:alpine3.16 as builder

WORKDIR /app/src

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o ./run ./cmd/main.go

FROM alpine:latest as production

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/src/run .

CMD ["./run"]