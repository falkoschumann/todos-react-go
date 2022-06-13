package portal

import "flag"

func ParseFlags() (host string, port uint) {
	flag.StringVar(&host, "host", "", "the server listen on this `host` (default all)")
	flag.UintVar(&port, "port", 8080, "the server listen on this `port`")
	flag.Parse()
	return
}
