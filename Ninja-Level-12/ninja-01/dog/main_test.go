package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(10)
	if y != 70 {
		t.Error("got", y, "expected", 70)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(10)
	if x != 70 {
		t.Error("got", x, "expected", 70)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	//Output:
	//70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	//Output:
	//70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}

}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}

}
