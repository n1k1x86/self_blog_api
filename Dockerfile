FROM golang:1.23-alpine as builder
WORKDIR /app
COPY ./api/go.mod ./api/go.sum ./
RUN go mod download

COPY ./api .

RUN CGO_ENABLED=0 GOOS=linux go build -o blog_api ./cmd
FROM alpine:latest

COPY --from=builder /app/blog_api /app/blog_api
COPY --from=builder /app/config/config.yaml /app/config.yaml
COPY --from=builder /app/db/migrations /app/migrations

WORKDIR /app
CMD [ "./blog_api" ]