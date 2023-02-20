FROM golang:latest-alpine

WORKDIR /usr/src/app 

COPY go.mod ./
COPY go.sum ./

RUN go mod tidy

COPY . .

EXPOSE 8090

CMD ["go", "run", "main.go"]