FROM alpine:3.17.3

ENV TZ Asia/Tokyo

RUN apk --update --no-cache add curl tzdata \
    && curl -sSf https://atlasgo.sh | sh

WORKDIR /root/migration

COPY . .

ENTRYPOINT ["/bin/sh"]
