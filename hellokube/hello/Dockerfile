# syntax=docker/dockerfile:1

FROM golang:1.16-buster AS build

WORKDIR /app

COPY / ./

RUN go mod init count
RUN go mod tidy

RUN go mod download
RUN go build -o main

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /app /app
WORKDIR /app

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["./main"]