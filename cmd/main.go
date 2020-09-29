package main


import (
	"flag"

	"github.com/2009_proto_h_server/pkg/server"
)


var (
	// Listenするアドレス+ポート
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	server.Serve(addr)
}

