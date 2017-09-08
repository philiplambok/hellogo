# Function

1. **sum is a function that takes a slice of numbers and adds them together. What
would its function signature look like in Go?**
	```go
	package main

	import "fmt"

	func main() {
		number := []int { 2, 3, 4 }
		fmt.Println(sum(number))
	}

	func sum(arr int) (total int) {
		total = 0
		for _, val := range arr {
			total += val
		}	
		return 
	}
	```

2. **Write a function that takes an integer and halves it and returns true if it was even or false if it was odd. For example, half(1) should return (0, false) and
half(2) should return (1, true) .**
	```go
	package main

	import "fmt"

	func main() {
		fmt.Println(half(1)) // return 0, false
		fmt.Println(half(2)) // return 1, true
	}

	func half(number int) (result int, status bool) {
		status = false
		result = 0
		
		if number%2 == 0 {
			status = true
		}

		result = number / 2

		return 
	}
	```

3. **Write a function with one variadic parameter that finds the greatest number in a
list of numbers**
	```go
	package main

	import "fmt"

	func main () {
		fmt.Println(greatest(2,3,4,212,421,222)) // return 421
	}	

	func greatest(args ...int) (number int) {
		number = 0

		for _, val := range args {
			if val > number {
				number = val
			}
		}

		return 
	}
	```

4. **Using makeEvenGenerator as an example, write a makeOddGenerator function
that generates odd numbers.**
	```go
	package main

	import "fmt"

	func main(){
		getOddNumber := generatorOddNumber()

		fmt.Println(getOddNumber()) // return 1
		fmt.Println(getOddNumber()) // return 3
		fmt.Println(getOddNumber()) // return 5
		fmt.Println(getOddNumber()) // return 7
	}	

	func generatorOddNumber() func() int {
		i := -1

		return func() int {
			i += 2
			return i
		}
	}
	```

5. **The Fibonacci sequence is defined as: fib(0) = 0 , fib(1) = 1 , fib(n) =
fib(n-1) + fib(n-2) . Write a recursive function that can find fib(n)**
	```go
	package main

	import "fmt"

	func main() {
		fmt.Println(fib(4))
	}

	func fib(number int) int {
		if number == 0 {
			return 0
		}

		if number == 1 {
			return 1
		}

		return fib(number-1) + fib(number-2)
	}
	```

6. **What are defer , panic , and recover ? How do you recover from a runtime panic?**

	Defer digunakan untuk menjalankan suatu fungsi di akhir ekseskusi suatu program. Biasanya pemakaiannya untuk memastikan suatu fungsi penutup file dijalankan diakhir setelah file telah dibaca. Panic untuk menghentikan suatu program, mirip fungsi `die()` di php. Namun pada Go parameter di dalam panic tidak keluar dengan sendirinya, untuk mengeluarkannya digunakan bantuan fungsi defer dan recover.

7. **How do you get the memory address of a variable?**
	
	Era pointer is backkk!! :))

	Pada golang kita dapat menggunakan pointer yang sudah ditingalkan banyak bahasa server, seperti php atau node.js. Untuk mengeluarkan alamat address(referensi) pada sebuah variabel ada dua cara, pertama jika anda mendeklarasikan variabel serperti biasa, *var number int* maka anda dapat mengeluarkan alamatnya dengan `&number`, namun jika anda menggunakan `number := new(int)`, maka anda dapat mengeluarkan alamatnya dengan `number`, dan valuenya dengan `*number`.

8. **How do you assign a value to a pointer?**

	untuk melalukan **asign** pada pointer dapat menggunakan `*pointer_var = value`

9. **How do you create a new pointer?**

	Untuk membuat sebuah variabel cukup mudah, hanya dengan `clone := &origin`, hanya memasukkan alamat referensinya saja. Untuk mengeluarkan nilai atau memodifikasi variabel pointer menggunakan `*clone = newValue`. Dengan memodifikasi variabel clone, maka variabel origin juga akan ikut berubah, karena mereka mengarah pada alamat memori yang sama.

10. **What is the value of x after running this program:**
	```go
	package main
	import "fmt"
	func square(x *float64) {
		*x = *x * *x
	}
	func main() {
		x := 1.5
		square(&x)
	}
	```
	Akan menghasilkan **2.25**. Untuk penggunaan variabel pointer dalam function, sebelumnya harus di deklarasikan dengan menggunakan *float64.

11. **Write a program that can swap two integers ( x := 1; y := 2; swap(&x, &y)
should give you x=2 and y=1 ).**
	```go
	package main
	import "fmt"
	func main() {
		a := 5 
		b := 4

		swap(&a, &b)

		fmt.Println(a) // return 4
		fmt.Println(b) // return 5
	}

	func swap(x, y *int) {
		temp := *x
		*x = *y
		*y = temp
	}
	``` 















