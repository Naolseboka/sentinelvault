FROM golang:1.22-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/team-service ./apps/team-service/cmd

FROM gcr.io/distroless/static-debian12
COPY --from=builder /bin/team-service /team-service
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/team-service"]
