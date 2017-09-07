# Arrays, Slice, Maps

1. **How do you access the fourth element of an array or slice?**

	Buat array terlebih dahulu.
	```go
	arr := [5]int {
		23,12,34,56,2
	}
	``` 
	akses array ke empat dengan `arr[3]`. Jika pada slice :
	```go
	arr := make([]int, 5){
		23,43,12,3,4,
	}
	```
	Akses array dengan `arr[3]`. 

2. **What is the length of a slice created using make([]int, 3, 9) ?**
	tiga (3)

3. **Write a program that finds the smallest number in this list:**
	```go
	x := []int{
	48,96,86,68,
	57,82,63,70,
	37,34,83,27,
	19,97, 9,17,
	}
	```
	Answer :
	```go
	package main

	import "fmt"

	func main() {

		x := []int{
			48, 96, 86, 68,
			57, 82, 63, 70,
			37, 34, 83, 27,
			19, 97, 9, 17,
		}
		max := 0

		for _, value := range x {

			if value > max {
				max = value
			}

		}

		fmt.Println(`Max : `, max)
	}
	```
