FROM golang:1.15.7-buster as builder

WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV port=8080

COPY . .

RUN go mod download
RUN go build cmd/api-grpc-web/main.go

FROM scratch
COPY --from=builder /app/ /app/

ENTRYPOINT ["/app/main"]

