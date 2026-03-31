FROM golang:1.25.8

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY generator ./generator

WORKDIR /app/generator
RUN go build -o generator

CMD ["./generator"]
