package main

import (
	"encoding/json"
	"fmt"
)

type singer struct {
	First string `json:"First"`
	Last  string `json:"Last"`
}

func main() {
	singers := `[{"First":"みな","Last":"名井"},{"First":"さな","Last":"湊崎"}]`
	bs := []byte(singers)

	var jp_singer []singer // jp_singer := []singer{}
	err := json.Unmarshal(bs, &jp_singer)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All the data: ", jp_singer)

	for i, v := range jp_singer {
		fmt.Println("\nSinger", i)
		fmt.Println(v.First, v.Last)
	}
	/*	fmt.Println(bs)
		fmt.Printf("%T\n", singers)
		fmt.Printf("%T\n", bs)
	*/

}
