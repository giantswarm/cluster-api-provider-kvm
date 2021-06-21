FROM alpine:3.14.0

RUN apk add --no-cache ca-certificates

ADD ./cluster-api-provider-kvm /cluster-api-provider-kvm

ENTRYPOINT ["/cluster-api-provider-kvm"]
