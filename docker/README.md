# Docker

## Docker commands

Build the backend for the app:

    `docker build --file docker/Backend.Dockerfile frontend/news-hack`
    `docker run -d -p 5001:5001 <container id>`

Build the frontend for the app:

    `docker build --file docker/Frontend.Dockerfile frontend/news-hack`
    `docker run -d -p 5173:5173 <container id>`

Build the Go API:

    `docker build --file docker/API.Dockerfile api -t news-hack-api`
    `docker run -p 8080:8080 news-hack-api`

Build all services using docker compose:

    docker compose -f docker/docker-compose.yml up --force-recreate
