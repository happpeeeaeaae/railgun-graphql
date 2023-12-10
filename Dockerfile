FROM golang:1.21.5-alpine as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /artifact/app ./main.go

FROM golang:1.21.5-alpine as artifact
WORKDIR /app

COPY --from=builder /artifact/app ./app

EXPOSE 8080
ENTRYPOINT ["./app"]