FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/gorilla/mux go.mongodb.org/mongo-driver/mongo
RUN go build -o main src/app/entry.go

CMD ["main"]

EXPOSE 8080