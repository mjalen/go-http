package route

import (
	"net"
	"fmt"

	"go-http/http/message"
)

type Router map[string]func(conn net.Conn)
type Request func(net.Conn)

func (r Router) Process(req []byte, conn net.Conn) string {
	ok := message.Validate("Simple-Request", req)	
	if !ok {
		fmt.Fprintf(conn, "400\r\n")
		return "400"
	}

	for p, callback := range r {
		if path := message.Find("Request-URI", req); string(path) == p {
			callback(conn)
			return "Provided route response."
		}
	}

	fmt.Fprintf(conn, "404\r\n")
	return "404"
}
