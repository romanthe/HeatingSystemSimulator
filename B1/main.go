// main.go
package main

import (
	conn "sandbox/connection"	
)

func main() {
	locker := make(chan bool, 1)	
	go conn.Server()
	go conn.Client(locker)

	// make infinite loop
	select{}	
}
