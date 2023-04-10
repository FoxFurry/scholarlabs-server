FROM golang:1.20-buster as builder

WORKDIR /app

COPY ./go.* ./
RUN go mod download

COPY . .

RUN cd services/course && go build -v -o server


FROM debian:buster-slim

COPY --from=builder /app/services/course/server /app/server

EXPOSE 22000

CMD ["/app/server"]
