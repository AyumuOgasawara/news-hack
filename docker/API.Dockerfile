FROM golang:1.23.3 AS go
WORKDIR /app
COPY go.mod go.sum main.go .env ./
RUN go mod download \
    && go build -o main /app/main.go
USER 1001
CMD [ "/app/main" ]
