FROM golang:1.25

WORKDIR /server

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o ./ .

RUN go build -o main .

CMD ["./main"]