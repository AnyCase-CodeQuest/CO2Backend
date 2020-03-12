#!/usr/bin/env bash
rm ./bin/main
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o bin/main src/app/entry.go
docker build -t coxa/co2backend:test .
docker push coxa/co2backend:test
kubectl rollout restart deployment/cobackend