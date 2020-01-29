## ENV
* DATABASE_CONNECTION = mongodb://some.url:123
* PASSWORD = password
* USERNAME = user
* DATABASE = dbname
* COLLECTION = some collection
### Test
* SEND_TO_QUEUE = 1
* QUEUE_NAME = sensors.myQueueName
* QUEUE_SERVER_HOST = host.name
* QUEUE_SERVER_PORT = 50000
## How to build and push image
* ```CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o bin/main src/app/entry.go```
* ```docker build -t coxa/co2backend:latest .```
* ```docker push coxa/co2backend:latest``` 