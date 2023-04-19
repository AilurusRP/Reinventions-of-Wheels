package main

import (
	"fmt"
	go_promise "go-promise-test/go-promise"
	"time"
)

func main() {
	go_promise.New(func(resolve func(any), reject func(string)) {
		fmt.Println("B")
		go func() {
			time.Sleep(time.Second * 3)
			resolve(1)
		}()
	}).Then(func(data any) any {
		fmt.Println("callback 1")
		fmt.Println(data)
		return 2
	}).Then(func(data any) any {
		fmt.Println("callback 2")
		fmt.Println(data)
		return 3
	}).Wait()
}
