# ---------- Build stage ----------
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server main.go


# ---------- Runtime stage ----------
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /app/server .

# NOTE: secrets are intentionally NOT baked into the image.
#   - app.env values are injected at runtime (compose env_file/environment).
#   - service-account.json is mounted read-only at runtime.
# See docker-compose.yml. This keeps credentials out of image layers and any
# registry the image is pushed to.

RUN mkdir -p uploads

EXPOSE 8080

CMD ["./server"]