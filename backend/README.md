# news-hack API

## RUN LOCALY

1. `go run main.go`

## DOCKER

1. `docker build ./ -t news-hac`
2. `docker run -p 8080:8080 news-hack`

## ENDPOINTS

- GET `/`
  get API description

- GET `/health`
  health check

- GET `/articles`
  get us news headlines

- POST `/articles`
  body {
  "keyword": string,
  "numArticles": number,
  "phoneNumber": string
  }
  NOTE: phoneNumber needs to have country code, you can hit this with curl command below

```sh
curl -X POST http://localhost:8080/articles \
    -H "Content-Type: application/json" \
    -d '{
        "keyword": "technology",
        "numArticles": 5,
        "phoneNumber": "441234567890"
    }'
```
