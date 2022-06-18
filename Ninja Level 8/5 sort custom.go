package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (bage ByAge) Len() int           { return len(bage) }
func (bage ByAge) Swap(i, j int)      { bage[i], bage[j] = bage[j], bage[i] }
func (bage ByAge) Less(i, j int) bool { return bage[i].Age < bage[j].Age }

type ByLast []user

func (blast ByLast) Len() int           { return len(blast) }
func (blast ByLast) Swap(i, j int)      { blast[i], blast[j] = blast[j], blast[i] }
func (blast ByLast) Less(i, j int) bool { return blast[i].Last < blast[j].Last }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	sort.Sort(ByAge(users))
	fmt.Println("Sort by Age")
	for _, v := range users {
		fmt.Println(v)
		sort.Strings(v.Sayings)
		for _, u := range v.Sayings {
			fmt.Println(u)
		}
	}

	sort.Sort(ByLast(users))
	fmt.Println("Sort by Last Name")
	for _, v2 := range users {
		fmt.Println(v2)
		sort.Strings(v2.Sayings)
		for _, u2 := range v2.Sayings {
			fmt.Println(u2)
		}
	}

}
