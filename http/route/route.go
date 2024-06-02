package route

import (
	"net"
	"slices"
	"fmt"

	"go-http/http/syntax"
)

type Router map[string]func(conn net.Conn)
type Request func(net.Conn)

func trim(req []byte) []byte {
	i := slices.Index(req, 0)
	return req[:i]
}

func (r Router) Process(req []byte, conn net.Conn) string {
	ok := syntax.Validate("Simple-Request", trim(req))	
	if !ok {
		fmt.Fprintf(conn, "400\r\n")
		return "400"
	}

	for p, callback := range r {
		if path := syntax.Find("Request-URI", req); string(path) == p {
			callback(conn)
			return "Provided route response."
		}
	}

	fmt.Fprintf(conn, "404\r\n")
	return "404"
}
