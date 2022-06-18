package main

import (
	"encoding/json"
	"fmt"
)

type Singer struct {
	Name  string
	Group string
}

func main() {
	var jsonBlob = []byte(`[
				{"Name": "Mina", "Group": "Twice"},
				{"Name": "Nancy", "Group": "Momoland"}
	]`)
	var singers []Singer
	err := json.Unmarshal(jsonBlob, &singers)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Printf("%+v", singers)
}
