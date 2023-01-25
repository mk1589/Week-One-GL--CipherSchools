package main

import "fmt"

func main() {
	//i := 0
	// for i < 3
	// for i := 0; i < 3; i++ {
	// 	if i == 1 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// 	//break
	// }
	//nums := []int{1, 3, 4, 8, 0}
	nums := []string{"manish",
		"mummy", "papa"}
	// for index, value := range nums
	// for _, value := range nums
	// for _, value := range "manish"
	for _, value := range nums {
		// if value == 4 {
		// 	// continue
		// 	break

		// }
		//fmt.Println(index)
		fmt.Println(value)
	}

}
