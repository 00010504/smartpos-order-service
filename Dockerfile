FROM golang:alpine3.16 as builder

RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init 
RUN  apk add librdkafka-dev pkgconf
RUN apk add make
#
RUN mkdir -p $GOPATH/src/gitlab.7i.uz/invan/invan_order_service 
WORKDIR $GOPATH/src/gitlab.7i.uz/invan/invan_order_service

# Copy the local package files to the container's workspace.
COPY . ./


# installing depends and build
RUN export CGO_ENABLED=1 && \
  export GOOS=linux && \
  make build && \
  mv ./bin/invan_order_service /


FROM surnet/alpine-wkhtmltopdf:3.16.2-0.12.6-small
RUN  apk add librdkafka-dev pkgconf
RUN wkhtmltopdf --version
COPY --from=builder invan_order_service .
ENTRYPOINT ["/invan_order_service"]
