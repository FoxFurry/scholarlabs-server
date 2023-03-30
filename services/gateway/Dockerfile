FROM golang:1.20-buster as builder

WORKDIR /app

COPY ./go.* ./
RUN go mod download

COPY ./gateway ./gateway

RUN cd gateway && go build -v -o server


FROM debian:buster-slim

COPY --from=builder /app/gateway/server /app/server

EXPOSE 8000

CMD ["/app/server"]
