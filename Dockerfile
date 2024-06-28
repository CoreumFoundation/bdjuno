FROM --platform=$TARGETPLATFORM alpine:latest
WORKDIR /bdjuno
RUN apk update
RUN apk add postgresql
COPY ./bin/.cache/bdjuno/docker.linux.amd64/bin/bdjuno /usr/bin/bdjuno
COPY database/schema /var/lib/postgresql/schema
RUN chmod a+rx /var/lib/postgresql && \
    chmod a+rx /var/lib/postgresql/schema

CMD [ "bdjuno", "start" ]
