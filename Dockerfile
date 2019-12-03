FROM golang:1.13

WORKDIR /go/src
COPY ./src /go/src

RUN go get -d -v github.com/gorilla/mux go.mongodb.org/mongo-driver/mongo
RUN GOOS=linux go build -o /main ./app/entry.go

CMD ["/main"]

EXPOSE 8080
