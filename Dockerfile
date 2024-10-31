FROM golang:1.23-bullseye

WORKDIR /wd

COPY go.mod .

RUN go mod download -x

COPY . .

CMD ["go", "run", "main.go"]
