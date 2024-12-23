FROM --platform=$TARGETPLATFORM alpine:latest
ARG TARGETOS
ARG TARGETARCH
WORKDIR /callisto
RUN apk update
RUN apk add postgresql
COPY ./bin/.cache/callisto/docker.$TARGETOS.$TARGETARCH/bin/callisto /usr/bin/callisto
COPY database/schema /var/lib/postgresql/schema
RUN chmod a+rx /var/lib/postgresql && \
    chmod a+rx /var/lib/postgresql/schema

CMD [ "callisto", "start" ]
