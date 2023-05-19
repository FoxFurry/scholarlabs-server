FROM golang:1.20-buster as builder

WORKDIR /app

COPY ./go.* ./
RUN go mod download

COPY . .

RUN cd services/environment && go build -v -o server


FROM debian:buster-slim

COPY --from=builder /app/services/environment/server /app/server

EXPOSE 22000

CMD ["/app/server"]
