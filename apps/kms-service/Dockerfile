FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/kms-service ./apps/kms-service/cmd

FROM gcr.io/distroless/static-debian12
COPY --from=builder /bin/kms-service /kms-service
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/kms-service"]
