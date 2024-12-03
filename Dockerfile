
FROM golang:1.23.3-alpine3.19 AS build

# Update repositories
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories \
    && apk update \
    && apk add --no-cache curl

RUN apk add --no-cache curl

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod

WORKDIR /app

COPY --from=build /app/main /app/main
COPY --from=build /app/frontend /app/frontend

EXPOSE ${PORT}
CMD ["./main"]

