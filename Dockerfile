FROM golang:1.12.1-alpine3.9
ENV GOPATH="/go"
RUN ["mkdir", "-p", "/go/src/github.com/rancher/demo"]
RUN ["mkdir", "-p", "/go/src/github.com/rancher/other"]
COPY main.go /go/src/github.com/rancher/demo/
COPY other.go /go/src/github.com/rancher/other/
WORKDIR /go/src/github.com/rancher/other
RUN ["go", "build", "-o", "other"]
WORKDIR /go/src/github.com/rancher/demo
RUN ["go", "build", "-o", "demo"]
ENTRYPOINT ["./other"]
