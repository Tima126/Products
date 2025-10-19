
FROM golang:1.25 AS builder


WORKDIR /app


COPY ./app/go.mod ./app/go.sum ./
RUN go mod download


COPY ./app/ ./


RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /usr/local/bin/server ./main.go


FROM debian:bullseye-slim

WORKDIR /app


COPY --from=builder /usr/local/bin/server ./


COPY ./app/templates ./templates


EXPOSE 8080

CMD ["./server"]
