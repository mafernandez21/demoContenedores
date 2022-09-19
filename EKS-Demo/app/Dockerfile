# docker build -t fortune-app:v1 --rm --pull --no-cache .
FROM golang:alpine3.16 as builder

WORKDIR /app

COPY . .

RUN go build -o fortune-app main.go

FROM alpine:3.16.2 as runner

WORKDIR /app

COPY --from=builder /app/public /app/public
COPY --from=builder /app/fortune-app /app/fortune-app
COPY --from=builder /app/fortune.txt /app/fortune.txt

EXPOSE 3000

CMD ["/app/fortune-app"]