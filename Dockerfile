FROM golang:1.10-alpine3.7
MAINTAINER Hung-Wei Chiu <hwchiu@linkernetworks.com>

WORKDIR /go/src/github.com/hwchiu/docker-multus-cni
COPY ./  /go/src/github.com/hwchiu/docker-multus-cni

RUN apk add --no-cache git bzr bash
RUN go get github.com/kardianos/govendor
RUN govendor sync
RUN git submodule init && git submodule update
WORKDIR multus-cni
RUN CGO_ENABLED=0 GOOS=linux ./build -a -tags netgo -ldflags '-w' && cp ./bin/multus /go/bin/multus

WORKDIR /go/src/github.com/hwchiu/docker-multus-cni
RUN go install .
ADD conf/ /tmp
ADD yaml/ /tmp
ADD entrypoint.sh /

ENV DEST_CNI /etc/cni/net.d/00-multus.conf
ENTRYPOINT ["/entrypoint.sh"]
