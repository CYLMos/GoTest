package main

// 引入套件
import "fmt"

// 主程式
func main() {
	value1 := 10

	fmt.Println(value1)

	if value1 < 11 {
		fmt.Println("value1 is smaller than 11")
	}

	for i := 0; i < 11; i++ {
		if i%2 == 1 {
			continue
		}

		fmt.Println(i)
	}

	var array1 [3]int64
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3

	array2 := [3]int64{4, 5, 6}

	for index, value := range array1 {
		fmt.Println(fmt.Sprintf("array1[%d]:%d", index, value))
	}

	for _, value := range array2 {
		fmt.Println(fmt.Sprintf("array2:%d", value))
	}

	// for i := 0; i < len(array1); i++ {
	// 	fmt.Println(fmt.Sprintf("array1[%d]:%d", i, array1[i]))
	// 	fmt.Println(fmt.Sprintf("array2[%d]:%d", i, array2[i]))
	// }
}
