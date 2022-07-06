package portal

import (
	"flag"
	"os"
	"strings"
)

const envPrefix = "TODOS_"

func ParseFlags() (host string, port uint, data string) {
	flag.StringVar(&host, "host", "", "the server is listening on this `host` (default all)")
	flag.UintVar(&port, "port", 8080, "the server is listening on this `port`")
	flag.StringVar(&data, "data", "todos.json", "todos saved in this `file`")

	flag.VisitAll(func(f *flag.Flag) {
		k := envPrefix + strings.ToUpper(f.Name)
		if f.Value.String() != f.DefValue {
			// CLI option overrides environment variable
			return
		}

		if v, ok := os.LookupEnv(k); ok {
			// Environment variable overrides default value of CLI option
			f.Value.Set(v)
		}
	})

	flag.Parse()

	return
}
