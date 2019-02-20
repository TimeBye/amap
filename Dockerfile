FROM golang:1.11.5-stretch
WORKDIR /go/src/github.com/TimeBye/amap
COPY . .
ENTRYPOINT ["/bin/bash","-c","go run main.go"]