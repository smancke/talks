## Golang im Docker Container

### Der einfache Weg: `golang:onbuild`

Dockerfile
```
FROM golang:1.6.0-onbuild
```

```bash
docker build -t google-query .
```

Aber, das resultierende image ist riesig:
```
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
google-query                latest              8d9199bf52f8        7 seconds ago       766.1 MB
```

### Minimal: Nur das GO binary

1. Statisch linken:
```bash
# Disable cgo
CGO_ENABLED=0 go build google_query.go
```

2. Docker images bauen
```
FROM scratch

COPY google_query /google_query
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

CMD ["/google_query"]
```

```
docker images
REPOSITORY                  TAG                 IMAGE ID            CREATED             SIZE
google-query-onbuild        latest              fbe6d751829b        2 minutes ago       766.3 MB
google-query-minimal        latest              c8cf922ef6f1        12 minutes ago      9.328 MB
```
