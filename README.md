# Social Network

A social network project with a React frontend and a Golang backend.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

## Prerequisites

- Node.js
- Go
- Docker
- Esbuild https://esbuild.github.io/

## <img src="https://cdn4.iconfinder.com/data/icons/logos-3/600/React.js_logo-512.png" width="28"> Running the React Frontend

After cloning the directory and changing into the `social-network` directory

- Run `npm install` to install dependencies
- Run `npm run dev` to start the development server

## <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_LightBlue.png" width="28"> Running the Golang Backend

After cloning the directory and changing into the `social-network` directory

- Run `go run .` to start the server
- Alternatively, run `go build main.go && ./main` to build and run the executable file

## <img src="https://www.docker.com/wp-content/uploads/2022/03/vertical-logo-monochromatic.png" width="28"> Using Docker

- Make sure you have the Docker daemon installed
  </br>
  To run the backend image using Docker, execute the following commands:
- Run $ `docker build -t social-network .` (The -t specifies a pseudo-tty or terminal for logging)
- Run $ `docker run -p 8080:8080 -it social-network` (The -p flag handles the port. The -it flags start a session with the container)
  Next, open another terminal for the frontend docker image and run the following:
- Run $ `docker build -t frontend -f frontend/Dockerfile .` (The -f specifies the location of the dockerfile)
- Run $ `docker run -p 8080:8080 -it social-network`

## Built With

- React
- Golang
- Docker
- HTML5
- Sqlite3

## Authors

LUIS316, jasonasante, gymlad, rw.smith
