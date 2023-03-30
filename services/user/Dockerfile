FROM golang:1.20-buster as builder

WORKDIR /app

COPY ./go.* ./
RUN go mod download

COPY ./services/user ./services/user

RUN cd services/user && go build -v -o server


FROM debian:buster-slim

COPY --from=builder /app/services/user/server /app/server

EXPOSE 8000

CMD ["/app/server"]
