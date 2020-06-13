FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o main .

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/main .


CMD ["./main", "start"] 