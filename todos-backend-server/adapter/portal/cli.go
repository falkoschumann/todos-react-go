package portal

import "flag"

func ParseFlags() (host string, port uint) {
	flag.StringVar(&host, "host", "", "the server is listening on this `host` (default all)")
	flag.UintVar(&port, "port", 8080, "the server is listening on this `port`")
	flag.Parse()
	return
}
