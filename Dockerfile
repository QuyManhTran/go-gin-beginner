FROM golang:1.24.1-alpine3.21 as builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev libc-dev sqlite-dev

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o main .

FROM golang:1.24.1-alpine3.21 as production

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]