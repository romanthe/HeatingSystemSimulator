package connection

import (
	mp "github.com/nkdm/HeatingSystemSimulator/B1/messageparser"
)

func registerRequest(port string) string {
	req := mp.Message{
		"id":     []string{MyID},
		"action": []string{"registerRequest"},
		"port":   []string{port},
	}
	return req.Encode()
}

func processRegisterResponse(s string) int {
	//resp := mp.ParseMessage(s)// need but only if i have to process info
	//unpack needed Value and call instruction
	// func foobar()
	//fmt.Println(resp["status"])// only for debug version
	return 0 // only for debug version
}
