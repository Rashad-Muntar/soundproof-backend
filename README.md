# Soundproof-backend

> This GitHub repository provides a robust API implementation for user authentication, utilizing JWT (JSON Web Tokens) for secure authentication and authorization. The API includes endpoints for user registration, login, token management (access token and refresh token), and user profile editing.

## Built With
- Golang
- Gin
- Gorm
- Posgresql
- Docker

## Getting Started

To set up a local version of this project, a collection of steps have been put together to help guide you from installations to usage. Simply follow the guide below and you'll be up and running in no time.

### Set up

- Install [Golang](https://go.dev/doc/install), if you haven't already.
- Install [Metamask](https://metamask.io/download/), if you haven't already.
- Install [Docker](https://docs.docker.com/), if you haven't already.
- Open Terminal
- Navigate to the preferred location/folder you want the app on your local machine. Use `cd <file-path>` for this.
- Run `git clone https://github.com/Rashad-Muntar/soundproof-backend.git` to download the source folder.
- Now that you have a local copy of the project, navigate to the root of the project folder from your terminal.
- Load the following .env keys 
  - `DB="host=localhost user=<db_user> password=<db_password dbname=<db_name> port=5432 sslmode=disable"`
  - `DB_NAME`
  - `DB_USER`
  - `DB_PASSWORD`
  - `JWT_SECRET=mncvbmcvnsdhmxmfkjghvmncxjghaduhgcmnbjbzfjgjkahjgdj`
- Run `docker build -t api-image .` to build and run the Docker containers
- Replace all environment variables in the docker-compose.yml file
- Run `docker-compose up -d`
- Open tool like Postman to test the enpoints below.

## Endpoints
_Base URL_: `http://localhost:8080`

|Description|Method|Endpoint|
|:---|:---|:---|
|Sign up to create an account|POST|`/auth/signup`|
|Log in to created account|POST|`/auth/login`|
|Fetch current user|GET|`/user/:userId`|
|Update user |PUT|`/user/:userId`|

### Usage

At this point, you now have everything you need to properly run the program (source code, node.js, mongodb, typescript). If not, refer back to the setup section of this documentation.

👤 **Rashad Muntar**

- GitHub: [@Rashad-Muntar](https://github.com/Rashad-Muntar)
- Twitter - [@RashadToure](https://twitter.com/RashadToure)
- LinkedIn: [Rashad Muntar](https://www.linkedin.com/in/rashad-muntar/)
