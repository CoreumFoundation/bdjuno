FROM --platform=$BUILDPLATFORM golang:1.18-alpine AS builder
RUN apk update && apk add --no-cache make git gcc libc-dev
WORKDIR /go/src/github.com/forbole/bdjuno
COPY . ./
ARG arch=x86_64
# we use the same arch in the CI as a workaround since we don't use the wasm in the indexer
ADD https://github.com/CosmWasm/wasmvm/releases/download/v1.1.1/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.${arch}.a
# Copy the library you want to the final location that will be found by the linker flag `-lwasmvm_muslc`
RUN cp /lib/libwasmvm_muslc.${arch}.a /lib/libwasmvm_muslc.a
RUN make build

FROM --platform=$TARGETPLATFORM alpine:latest
COPY --from=builder /go/src/github.com/forbole/bdjuno/build/bdjuno /usr/bin/bdjuno
WORKDIR /bdjuno
RUN apk update
RUN apk add postgresql sudo
RUN adduser -DG wheel bdjuno
RUN sed -e 's;^# \(%wheel.*NOPASSWD.*\);\1;g' -i /etc/sudoers
USER bdjuno
COPY --chown=bdjuno database/schema /var/lib/postgresql/schema

CMD [ "bdjuno", "start" ]
