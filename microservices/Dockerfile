## Build App
FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY . ./

RUN apk add --no-cache git
RUN go mod tidy
# RUN go install github.com/swaggo/swag/cmd/swag@latest

# RUN swag init -g api.go
RUN go build -o /docker-app main.go

# EXPOSE 8080

# CMD [ "/docker-app" ]

## Deploy
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /docker-app .

ENV PORT 9090
EXPOSE 9090

CMD ["./docker-app"]
