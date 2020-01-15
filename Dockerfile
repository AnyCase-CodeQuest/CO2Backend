FROM golang:1.13

WORKDIR /app
COPY ./src /app/src
COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum
RUN GOOS=linux GOARM=7 go build -o /app/main /app/src/app/entry.go
COPY ./build /app/build

CMD ["/app/main"]

EXPOSE 8080
