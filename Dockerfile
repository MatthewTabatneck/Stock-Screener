# syntax=docker/dockerfile:1

# ---- Build stage ----
FROM golang:1.25.1 AS build
WORKDIR /app

# Copy go.mod first and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy all source code
COPY . .

# Build two binaries
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./cmd/api && \
    CGO_ENABLED=0 GOOS=linux go build -o /bin/worker ./cmd/worker

# ---- Runtime stage ----
FROM gcr.io/distroless/base-debian12
ENV PORT=8080
COPY --from=build /bin/api /bin/api
COPY --from=build /bin/worker /bin/worker

# Default command: API service
CMD ["/bin/api"]
