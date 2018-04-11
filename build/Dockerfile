FROM alpine

ENV RABBITMQ_SERVER=rabbitmq-rabbitmq \
    RABBITMQ_USER=user \
    RABBITMQ_PASS=arJNaFrszk \
    POSTGRESQL_ARGS="host=postgresql-postgresql port=5432 user=postgres dbname=postgres password=qBDXNlz276 sslmode=disable"

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY sitemgr /

ENTRYPOINT /sitemgr

