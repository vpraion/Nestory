# Stage 1 : build
FROM golang:tip-alpine3.24 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o server main.go

# Stage 2 : runner
FROM golang:tip-alpine3.24 AS runner

RUN adduser -D -u 10001 appuser

WORKDIR /app

RUN chown -R appuser:appuser /app

COPY --from=builder --chown=appuser:appuser /app/server .

USER appuser

ARG APP_PORT=8080
ENV APP_PORT=${APP_PORT}

EXPOSE ${APP_PORT}

HEALTHCHECK --start-period=1s --interval=2s --timeout=3s --retries=3 CMD wget -qO- http://localhost:${APP_PORT}/healthz || exit 1

CMD ["./server"]




