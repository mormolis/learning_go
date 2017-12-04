// go build -o fizzbuzz fizzbuzz
package fizzbuzz

import (
	"fmt"
	"math"
)

func FizzBuzz(num int) string {
	if math.Mod(float64(num), 15) == 0 {
		return "FizzBuzz"
	} else if math.Mod(float64(num), 5) == 0 {
		return "Buzz"
	} else if math.Mod(float64(num), 3) == 0 {
		return "Fizz"
	}
	return fmt.Sprintf("%d", num)
}

func PlayFizzBuzz() {
	for i := 0; i < 200; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
