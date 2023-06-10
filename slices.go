package main

import "fmt"

func main() {

	var emptySlice []string
	fmt.Println("Empty Slice:", emptySlice)
	fmt.Println("Length of Empty Slice:", len(emptySlice))
	fmt.Println("Is empty slice equal to nil:", emptySlice==nil)

	exampleArray := [5]int{2, 3, 4, 5, 6}
	fmt.Println("Example Array:", exampleArray)
	fmt.Println("Example Array Length:", len(exampleArray))
	fmt.Println("Example Array Capacity:", cap(exampleArray))

	exampleSlice := exampleArray[:3]
	fmt.Println("Example Slice:", exampleSlice)
	fmt.Println("Example Slice Length:", len(exampleSlice))
	fmt.Println("Example Slice Capacity:", cap(exampleSlice))

	exampleSlice = append(exampleSlice,2)
	exampleSlice = append(exampleSlice,2)
	exampleSlice = append(exampleSlice,2)
	fmt.Println("Example Slice:", exampleSlice)
	fmt.Println("Example Array:", exampleArray)
	fmt.Println("Example Array Length:", len(exampleArray))

	exampleSliceCopied := make([]int, len(exampleSlice))
	fmt.Println("Example Copied Empty:", exampleSliceCopied)
	copy(exampleSliceCopied, exampleSlice)
	fmt.Println("Example Copied:", exampleSliceCopied)

	twoD := make([][]int, len(exampleArray))
	for i := 0; i<len(exampleArray); i++{
		twoD[i] = make([]int, len(exampleArray))
		for j := 0; j<len(exampleArray); j++{
			twoD[i][j]=i+j
		}
	}
	fmt.Println("Two D:", twoD)
}