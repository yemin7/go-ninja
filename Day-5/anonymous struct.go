package main

import "fmt"

func main() {

	jp := struct {
		first, last string
		age         int
	}{
		first: "みな",
		last:  "名井",
		age:   26,
	}

	fmt.Println(jp.first, jp.last, jp.age)
}
