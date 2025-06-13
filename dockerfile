FROM golang:1.23.4-alpine3.20 AS builder

ENV GO111MODULE=on
ENV CGO_ENABLED=0

WORKDIR /app

COPY cmd ./cmd
COPY internal ./internal
COPY go.mod go.sum main.go ./

RUN go mod tidy
RUN go build -o /go/bin/ossinsight-mcp

FROM alpine:latest

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

WORKDIR /app

COPY --from=builder /go/bin/ossinsight-mcp /app/ossinsight-mcp

EXPOSE 8080

# CMD ["/app/ossinsight-mcp", "server", "--host", "0.0.0.0", "--port", "8080"]
# CMD ["/app/ossinsight-mcp", "client", "--endpoint", "http://localhost:8080/mcp"]

# docker buildx build --platform linux/amd64 -t ossinsight-mcp:latest .
# docker run -d -p 8080:8080 --env-file .env --name ossinsight-mcp -t ossinsight-mcp:latest /app/ossinsight-mcp server --host 0.0.0.0 --port 8080