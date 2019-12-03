FROM golang:1.13

RUN go get -d -v github.com/gorilla/mux go.mongodb.org/mongo-driver/mongo

WORKDIR /app
COPY ./src /go/src
COPY ./build /app/build

RUN GOOS=linux GOARM=7 go build -o /app/main /go/src/app/entry.go

CMD ["/app/main"]

EXPOSE 8080
