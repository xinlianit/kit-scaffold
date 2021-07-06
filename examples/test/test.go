package main

import "fmt"

func main(){
	array1 := []int{1,2,3,4,5}
	array2 := []int{4,5,6,7,8}
	arrayMerge := arrayMerge(array1, array2)
	fmt.Println("合并后数组：", arrayMerge)
}

func arrayMerge(array1 []int, array2 []int) []int {
	var result []int

	for k, v := range array1 {
		fmt.Printf("key: %d, val: %d \n", k, v)
		result = append(result, v)
	}

	for k, v := range array2 {
		fmt.Printf("key: %d, val: %d \n", k, v)
		result = append(result, v)
	}

	return result
}


