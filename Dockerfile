FROM alpine:latest

ADD multus-cni/bin/multus /tmp
ADD conf/ /tmp
ADD yaml/ /tmp
ADD entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]
