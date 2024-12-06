
FROM golang:1.23.3-alpine3.19 AS build

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

