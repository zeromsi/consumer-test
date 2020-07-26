FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go env -w GOPROXY="https://goproxy.io,direct"
RUN go mod download
COPY . .
RUN go build -o /app/bin/consumer .


FROM golang:latest
WORKDIR /app
COPY --from=builder /app/bin /app
EXPOSE 8080
CMD ["./consumer"]