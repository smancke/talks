
FROM alpine

RUN apk --update add curl git \
  && curl https://caddyserver.com/download/linux/amd64?plugins=http.jwt,http.login,http.git > /caddy.tar.gz \
  && tar xvfz /caddy.tar.gz \
  && chmod a+x /caddy

ADD Caddyfile /

CMD ["/caddy", "-conf", "/Caddyfile"]

