FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/access-request-service ./apps/access-request-service/cmd

FROM gcr.io/distroless/static-debian12
COPY --from=builder /bin/access-request-service /access-request-service
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/access-request-service"]
