package connection

import (
	"fmt"
	mp "sandbox/messageparser"
)

//TODO for now controller calls request functions staticly,
//TODO dynamic method is planed
func requestController(s string) string {
	request := mp.ParseMessage(s)
	fmt.Println(request) // temp instruction to erase in final version	
	action, ok := request["action"]
	if !ok {
		fmt.Println("Request Controller Error!")
	}
	switch action[0] {
	case "getDataRequest":
		return getDataRequest(request)
	case "setControllerRequest":
		return setControllerRequest(request)
	default:
		//unknown request
	}
	return "id name action=Response error=unknownRequest"
}

func getDataRequest(m mp.Message) string {
	//unpack needed Value and call function
	// func foobar()
	resp := mp.Message{
		"id":     []string{"name"},
		"action": []string{"getDataResponse"},
		"th":     []string{"40.00"},
		"tr":     []string{"20.07"},
		"tpco":   []string{"62.33"},
		"p":      []string{"2"},
		"i":      []string{"0.34"},
		"d":      []string{"0.444"},
	}
	return resp.Encode()
}

func setControllerRequest(m mp.Message) string {
	//unpack needed Value and call function
	// func foobar()
	resp := mp.Message{
		"id":     []string{"name"},
		"action": []string{"setControllerResponse"},
		"status": []string{"0"},
	}
	return resp.Encode()
}
