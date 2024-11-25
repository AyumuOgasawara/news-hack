# news-hack
ニュースを指定した電話番号のWhatsAppに送信するアプリです。ユーザーは、`キーワード`、`ニュースの件数`、`送信先の電話番号`を入力します。すると送信先の電話番号に指定したキーワードの入ったニュースが指定した件数分送信されます。


## なんのために作られたか
ニュースを読む習慣をつけるために受動的にニュースに触れることができたらいいと思い作成しました。<br>
現段階はプロトタイプで、今後指定した電話番号に指定した頻度と時間で自動的にニュースが送信されるようにしたいと考えています。

## Architecture Diagram
![news-hack](https://github.com/user-attachments/assets/a8f71076-1d4f-4120-af12-fb2b7bb7a966)

## Technology Stack

| **Category**       | **Technology**                 | **Version**          |
|--------------------|---------------------------------|----------------------|
| **APP**       | [React](https://react.dev/)                          | 18.2.0              |
|        | [Node.js](https://nodejs.org/en)                          |  v22.9.0        |
| **API**        | [Go](https://go.dev/)                             | 1.23.3              |
| **API Framework** | [Gin Gonic](https://github.com/gin-gonic/gin)                   | 1.10.0              |
| **Environment Management** | [godotenv](https://github.com/joho/godotenv)             | 1.5.1               |

## 必要条件
### 前提条件
Node.jsとGoのインストールが終了している。

## 外部サービスの設定
**NEWS API**
- 公式に従って登録し、API KEYを取得してください。

**WhatsApp**
- 公式に従って登録し、アクセストークンを取得してください。

### 環境変数
以下は`api/`下で`.env`ファイルを作成し、記載してください。
| 環境変数 | 用途 |
| ------------ | ------- |
| NEWS_API_KEY | NEWS APIのAPI KEY|
| WHATSAPP_TOKEN | WhatsAppのトークン(時間制限有) |


## How to run
### Running Locally
#### APP README
https://github.com/AyumuOgasawara/news-hack/blob/main/frontend/news-hack/README.md

#### API README
https://github.com/AyumuOgasawara/news-hack/blob/main/api/README.md

### Using Docker
https://github.com/AyumuOgasawara/news-hack/blob/main/docker/README.md<br>
* `docker compose` doesn't work so, if you want to use `Docker`, run them separately.