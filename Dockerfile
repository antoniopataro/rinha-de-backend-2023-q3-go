FROM golang:1.20.6 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o ./main ./cmd/main.go

EXPOSE 8080

FROM alpine:3.14.10

COPY --from=builder /app/main /main

CMD ["/main"]