# Stage 1 : build
FROM golang:tip-alpine3.24 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Stage 2 : runner
FROM golang:tip-alpine3.24 AS runner

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]




