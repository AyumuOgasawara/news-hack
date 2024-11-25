# news-hack
This application sends news articles to a specified phone number via WhatsApp. Users can input a `keyword`, the `number of news articles`, and the `recipient's phone number`. The application will then send the specified number of news articles containing the given keyword to the designated phone number.

## Purpose
This app was developed to encourage a habit of reading news by providing a passive way to engage with news articles.  
Currently, it is in the prototype stage. In the future, the app aims to automatically send news to specified phone numbers at customized frequencies and times.

## Architecture Diagram
![news-hack](https://github.com/user-attachments/assets/a8f71076-1d4f-4120-af12-fb2b7bb7a966)

## Technology Stack

| **Category**           | **Technology**                 | **Version**          |
|-------------------------|---------------------------------|----------------------|
| **Frontend (APP)**      | [React](https://react.dev/)    | 18.2.0              |
|                         | [Node.js](https://nodejs.org/) | 22.9.0              |
| **Backend (API)**       | [Go](https://go.dev/)          | 1.23.3              |
| **API Framework**       | [Gin Gonic](https://github.com/gin-gonic/gin) | 1.10.0              |
| **Environment Manager** | [godotenv](https://github.com/joho/godotenv)  | 1.5.1               |

## Prerequisites
### Installation
Ensure the following are installed on your system:
- [Node.js](https://nodejs.org/en)
- [Go](https://go.dev/)

## External Service Setup
### NEWS API
- Follow the instructions in [the official NEWS API documentation](https://newsapi.org/docs/get-started) to register and obtain an API key.

### WhatsApp
- Follow the instructions in [the official WhatsApp documentation](https://developers.facebook.com/docs/whatsapp/) to register and obtain an access token.

### Environment Variables
Create a `.env` file under the `api/` directory with the following content:

| Variable         | Description                          |
|-------------------|--------------------------------------|
| `NEWS_API_KEY`    | API key for the NEWS API            |
| `WHATSAPP_TOKEN`  | WhatsApp access token (time-limited)|

## How to Run

### Running Locally
#### Frontend
See the [Frontend README](https://github.com/AyumuOgasawara/news-hack/blob/main/frontend/news-hack/README.md).

#### Backend
See the [API README](https://github.com/AyumuOgasawara/news-hack/blob/main/api/README.md).

### Using Docker
Refer to the [Docker README](https://github.com/AyumuOgasawara/news-hack/blob/main/docker/README.md).  
**Note**: `docker compose` is currently not functional. To use Docker, run the services individually as described in the documentation.