FROM golang:1.8-alpine

ADD main.go /go/src/content/
RUN apk --no-cache add git curl \
    && go get -v content

EXPOSE 80

CMD ["content"]

HEALTHCHECK --interval=3s \
  CMD curl -f http://localhost/ || exit 1
  