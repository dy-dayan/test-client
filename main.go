package main

import (
	"flag"
	"github.com/dy-dayan/test-client/client"
	"github.com/dy-dayan/test-client/unit-test"
)




func main() {
	var addr string
	flag.StringVar(&addr, "d", "0.0.0.0:8080", "remote host address")
	flag.Parse()
	c := client.NewClient(client.Addr(addr))

	go c.GetMessage()
	for {
		unitTest.Hello(c)
	}
}



