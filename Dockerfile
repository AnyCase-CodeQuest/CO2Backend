FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/gorilla/mux
RUN go build main.go

CMD ["./main"]

EXPOSE 8080
