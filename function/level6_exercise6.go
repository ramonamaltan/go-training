package main

import "fmt"

// anonymous self executing function
func main() {
	func() {
		fmt.Println("anonymous")
	}()
}
