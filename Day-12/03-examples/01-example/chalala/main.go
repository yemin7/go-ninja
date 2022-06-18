//Package chalala asks if you are ready to rock
package chalala

//Sum adds an unlimited number of values of type int
func Sumf(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
