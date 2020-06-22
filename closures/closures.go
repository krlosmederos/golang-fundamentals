package closures

import "fmt"

func main() {
	var val int

	fmt.Scanf("%d", &val)

	myFunc := Product(val)

	for i := 0; i < 10; i++ {
		fmt.Println(myFunc())
	}
}

// Product is a function that return other function
func Product(num int) func() int {
	counter := 0

	return func() int {
		counter++
		return num * counter
	}
}
