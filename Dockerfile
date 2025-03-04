FROM golang:1.23.6

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd/api/main.go

EXPOSE 8080

CMD [ "./app" ]

