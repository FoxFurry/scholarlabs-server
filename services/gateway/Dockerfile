FROM golang:1.20-buster as builder

WORKDIR /app

COPY ./go.* ./
RUN go mod download

COPY ./services/gateway ./services/gateway

RUN cd services/gateway && go build -v -o server


FROM debian:buster-slim

COPY --from=builder /app/services/gateway/server /app/server

EXPOSE 8000

CMD ["/app/server"]
