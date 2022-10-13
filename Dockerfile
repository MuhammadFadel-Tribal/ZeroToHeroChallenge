FROM golang:1.18.4-alpine AS build_base

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download && go mod verify

COPY . .
RUN cd cmd && go build -o ./out

EXPOSE 8080

CMD ["./cmd/out"]
