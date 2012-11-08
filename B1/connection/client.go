// Connection handler package.
package connection

import (
	"fmt"
	"net"
	"time"
)

// Configure client and start infinite time trigered requests
// TODO Add port and rest parameters
func Client(sem chan bool) {
	for {
		DialDLS(sem, UsePort)
		<-time.After(HelloTime * time.Second)
	}

}

//TODO change channel semaphore to be more efficient
func DialDLS(sem chan bool, port string) {

	sem <- true
	tcpAddr, err := net.ResolveTCPAddr("tcp4", DLSIp+":"+UsePort)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	var buf [512]byte

	s := registerRequest(port)
	fmt.Println("Send request:", s)

	_, err1 := conn.Write([]byte(s))
	if err1 != nil {
		return
	}

	n, err2 := conn.Read(buf[0:])
	if err2 != nil {
		return
	}
	<-sem

	fmt.Println("Recieve response: ", string(buf[0:n]))
	processRegisterResponse(string(buf[0:n]))
}
