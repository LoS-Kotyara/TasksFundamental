FROM golang:1.17.6 AS builder

RUN mkdir /app
RUN git clone https://github.com/LoS-Kotyara/TasksFundamental.git /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -o dist/main

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./dist/main"]
