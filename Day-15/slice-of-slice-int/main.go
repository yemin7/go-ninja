package main

import "fmt"

func main() {
	s2s := make([][]string, 0, 4)

	s1 := make([]string, 3, 3)
	s1[0] = "a"
	s1[1] = "b"
	s1[2] = "c"

	s2 := make([]string, 3, 3)
	s2[0] = "a2"
	s2[1] = "b2"
	s2[2] = "c2"

	//s2s = append(s2s, s1)
	//fmt.Println("Append s1 to s2s", s2s)

	s1 = append(s1, "d")
	s2s = append(s2s, s1)
	fmt.Print(cap(s1))
	fmt.Println("Append s1 to s2s", s2s)

	s2s = append(s2s, s2)
	fmt.Println("Append s2 to s2s", s2s)

	a1 := [4]string{"A", "B", "C", "D"}
	student := make([]string, 4)
	for i := 0; i < 4; i++ {
		student[i] = a1[i]
	}
	fmt.Println(student)
}
