FROM golang:1.6

ENV PROJECT_PATH=/go/src/github.com/Magicking/ether-swarm-cache/

COPY . $PROJECT_PATH
WORKDIR $PROJECT_PATH

RUN go build -v -o /ether-swarm-cache main.go

EXPOSE 8091

ENTRYPOINT ["/ether-swarm-cache", "--host", "0.0.0.0", "--port", "8091"]
