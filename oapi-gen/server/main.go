// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml

package main

import (
	"flag"
	"server/swsrv"
)

func main() {
	var port = flag.Int("port", 5001, "Port for test HTTP server")
	flag.Parse()

	swsrv.StartServer(*port)
}
