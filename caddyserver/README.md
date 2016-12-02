
# Caddyserver

Serve The Web Like It's 2016

Caddy is an alternative web server that is easy to configure and use.

[caddyserver.com](http://caddyserver.com/)

## Installation

### Way one: The golang way
```shell
go get  github.com/mholt/caddy/caddy
```

### Way two: Custom Caddy
```go
package main

import (
    "github.com/mholt/caddy/caddy/caddymain"

    _ "github.com/BTBurke/caddy-jwt"
    _ "github.com/abiosoft/caddy-git"
    _ "github.com/tarent/loginsrv/caddy"
)

func main() {
        caddymain.Run()
}
```

### Way three: Click & Download

https://caddyserver.com/download


## Caddy Hello World

Just start to serve the local directory:
```shell
./caddy
```

## Features

### TLS and Let's Encrypt
Caddy has build in support for automatic TLS
and of course serves HTTP2.

E.g. Caddyfile with a self signed certificate:
```
https://localhost:8080/

tls self_signed
```

### Git
Automated Git Checkout
```
git git@git.mancke.net:mancke.net/web.git {
                key      /cfg/id_rsa
                path     /tmp/caddy-demo-data
                hook /git-trigger secret
}
```

### Marktdown support


## Securing with JWT and Osiam

* [JWT Plugin](https://github.com/BTBurke/caddy-jwt)
* [Loginsrv](https://github.com/tarent/loginsrv)
* [OSIAM](https://osiam.org/)

```
http://localhost:2015 {

  jwt {
    path /
    redirect /login
  }
  
  loginsrv / {
    success-url /
    backend provider=simple,bob=secret
    backend provider=osiam,endpoint=http://localhost:8080,clientId=example-client,clientSecret=secret
  }

}
```


## UI Composition

[caddy-uic](https://github.com/smancke/caddy-uic/)
[lib-compose](https://github.com/tarent/lib-compose/blob/master/composition/README.md)

