package main

import (
	//_ "github.com/BTBurke/caddy-jwt"
	_ "github.com/abiosoft/caddy-git"
	"github.com/mholt/caddy/caddy/caddymain"
	_ "github.com/smancke/caddy-jwt"
	_ "github.com/smancke/caddy-uic"
	_ "github.com/tarent/loginsrv/caddy"
	_ "github.com/tarent/loginsrv/htpasswd"
	_ "github.com/tarent/loginsrv/osiam"
)

func main() {
	caddymain.Run()
}
