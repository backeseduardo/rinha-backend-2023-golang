FROM golang:1.19

WORKDIR /app

#COPY go.mod go.sum ./
COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /rinha-dev cmd/web/main.go

CMD ["/rinha-dev"]

