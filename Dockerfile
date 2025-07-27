FROM golang:1.22.3 AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o server

FROM scratch
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]

