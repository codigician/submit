# build stage
FROM golang:1.17

WORKDIR /app

COPY . /app

ENV CGO_ENABLED=0

RUN go build -o submit-api .

# execution stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY --from=0 /app/submit-api ./

EXPOSE 8084

CMD ["./submit-api"]
