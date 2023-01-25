package main

import "fmt"

func main() {
	result, x := c()
	fmt.Println(result)
	fmt.Println(x)

}

func a() (int, string) {
	return 15, "manish"

}
func b() {

}
func c() (i int, j string) {
	i = 10
	j = "deepak"
	return

}
