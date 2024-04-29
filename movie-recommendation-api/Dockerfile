FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY /movie-recommendation-api/go.mod /movie-recommendation-api/go.sum ./

RUN go mod download

COPY /movie-recommendation-api/. .

RUN go build -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=build /app/main .

CMD ["./main"]
