package main

import (
	"fmt"
	"nowgoal/protocol"
)

func main() {
	err := protocol.ServeHTTP()
	if err != nil {
		result := fmt.Errorf("Error : %v", err)
		fmt.Println(result)
	}
}
