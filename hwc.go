package main 

import "fmt"

func main() {
	done := make(chan struct{})
	go func() {
		fmt.Println("Hello Concurrent World!")
		close(done)
	}()
	<-done
}
