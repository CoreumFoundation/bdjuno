FROM --platform=$BUILDPLATFORM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache make git gcc libc-dev
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./
ARG arch=x86_64
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.${arch}.a
# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a
RUN make build

FROM --platform=$TARGETPLATFORM alpine:latest
WORKDIR /bdjuno
RUN apk update
RUN apk add postgresql
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
COPY database/schema /var/lib/postgresql/schema

CMD [ "bdjuno", "start" ]
