FROM --platform=$BUILDPLATFORM golang:1.21-bullseye AS builder
RUN apt update && apt install make git gcc libc-dev
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./
RUN go mod download
RUN make build

FROM --platform=$TARGETPLATFORM alpine:latest
WORKDIR /bdjuno
RUN apk update
RUN apk add postgresql
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno/bdjuno /usr/bin/bdjuno
COPY database/schema /var/lib/postgresql/schema
RUN chmod a+rx /var/lib/postgresql && \
    chmod a+rx /var/lib/postgresql/schema

CMD [ "bdjuno", "start" ]
