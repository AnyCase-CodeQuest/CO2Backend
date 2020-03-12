FROM alpine:3.11

WORKDIR /app
COPY ./bin/main /app/main

CMD ["/app/main"]

EXPOSE 8080
