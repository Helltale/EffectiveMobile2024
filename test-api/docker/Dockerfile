FROM golang:1.22.4 AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./

RUN go mod download

COPY . .

RUN go build -o main main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/main .

WORKDIR /

CMD ["./main"]