package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var result int
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		time.Sleep(2 * time.Millisecond)
		result = 42
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(result)
}
