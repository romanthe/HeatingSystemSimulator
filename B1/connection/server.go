// Connection handler package.
package connection

import (
	"fmt"
	"net"
	"os"
)

// Configure server and start infinite listening loop
// TODO Add port and rest parameters
func Server() {

	listener, err := net.Listen("tcp", ":" + UsePort)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte

	n, err := conn.Read(buf[0:])
	if err != nil {
		return
	}

	s := requestController(string(buf[0:n]))

	_, err2 := conn.Write([]byte(s))
	if err2 != nil {
		return
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
