# syntax=docker/dockerfile:1
# ---- Build Stage ----
FROM golang:1.23 AS builder
WORKDIR /app

# Enable auto toolchain if deps ask for a newer point release
ENV CGO_ENABLED=0 GOOS=linux GOTOOLCHAIN=auto

# Copy mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest
COPY . .

# Build the binary
RUN go build -o server ./cmd/server

# ---- Run Stage ----
FROM gcr.io/distroless/base-debian12 AS RUNNER
WORKDIR /
ENV PORT=3001 \
    UPLOAD_DIR=/uploads
COPY --from=builder /app/server /server
COPY --from=builder /app/docs /docs
VOLUME ["/uploads"]
EXPOSE 3001
ENTRYPOINT ["/server"]
