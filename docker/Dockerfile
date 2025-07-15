FROM golang:1.23.4-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app_binary ./src/cmd/main.go

FROM alpine:latest

COPY --from=builder /app_binary /app_binary

COPY templates/ /templates/

EXPOSE 8050

CMD ["/app_binary"]