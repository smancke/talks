
FROM alpine

RUN apk --update add curl \
  && curl https://caddyserver.com/download/linux/amd64?plugins=http.jwt,http.login > /caddy.tar.gz \
  && tar xvfz /caddy.tar.gz \
  && chmod a+x /caddy

ADD Caddyfile /
ADD index.html /webroot/index.html

CMD ["/caddy", "-conf", "/Caddyfile"]

