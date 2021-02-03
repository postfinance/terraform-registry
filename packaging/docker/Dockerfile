FROM alpine:latest

ENV TZ=Europe/Zurich


COPY registry /usr/bin/registry
COPY packaging/docker/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod 750 /docker-entrypoint.sh

CMD ["--help"]
ENTRYPOINT ["/docker-entrypoint.sh"]
