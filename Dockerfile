FROM golang:alpine AS builder

RUN apk update \
    && apk add git \
    && apk add build-base \
    && apk add pcre-dev \
    && apk add sqlite-dev \
    && apk --no-cache add build-base

WORKDIR /app

COPY . .

RUN make build

RUN pwd
RUN ls

FROM alpine:latest

RUN apk update \
    && apk add pcre \
    && apk add sqlite

COPY --from=builder /app/bin/phone-validator ./
COPY --from=builder /app/sample.db ./
COPY --from=builder /app/sqlite3_mod_regexp.so /usr/local/lib/
RUN ls
CMD [ "/phone-validator" ]