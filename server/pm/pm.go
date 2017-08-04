package pm

import (
	"fmt"
	"math"
)

func add(x int, y [2]int) int {
	return x + y[0] + y[1]
}

func del(x int, y int) int {
	return x - y
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func printFor() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func Sqrt(x float64) float64 {
	z := float64(1)
	for {
		tmp := z - (z*z-x)/(2*z)
		fmt.Println(tmp)
		if tmp == z || math.Abs(tmp-z) < 0.0000000000001 {
			break
		}
		z = tmp
	}
	return z
}

func Init() {

}
