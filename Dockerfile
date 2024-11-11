FROM golang:1.21.0-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o /quangdz

ENTRYPOINT ["/quangdz"]
