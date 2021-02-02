FROM linux-docker-local.repo.pnet.ch/pf/alpine:latest

ENV TZ=Europe/Zurich

RUN http_proxy=http://outapplpf.pnet.ch:3128 apk add --no-cache tzdata


COPY cmd /usr/bin/cmd
COPY packaging/docker/docker-entrypoint.sh /docker-entrypoint.sh
RUN chmod 750 /docker-entrypoint.sh

CMD ["--help"]
ENTRYPOINT ["/docker-entrypoint.sh"]
