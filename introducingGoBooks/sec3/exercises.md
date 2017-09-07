# Variables

1. **What are two ways to create a new variable?**
	Pada golang ada dua cara dalam mendeklarasikan sebuah variabel.
	
	- Menggunakan kata kunci `var`
		```go
		package main

		import "fmt"

		func main() {
			// Versi disambung
			var name string = "Philip Lambok"
			fmt.Println(name)

			// Versi dipisah
			var name string
			name = "Philip Lambok"
		}
		```
	- Tanpa menggunakan kata kunci `var`. Dimana penggunaan inilah yang dianjurkan atau standarisasi dari golang. 
		```go
		package main

		import "fmt"

		func main() {
			name := "Philip Lambok"
			fmt.Println(name)
		}
		```

2. **What is the value of x after running x := 5; x += 1**

	6 (konsep variable)

3. **What is scope? How do you determine the scope of a variable in Go?**

	Scope adalah aturan akses sesuatu pada sebuah pemrograman. Pada Go Scope diatur oleh `{ }`. Dimana Variable yang diatas(external) dapat diakses oleh scope yang berada di dalamnya, tidak sebaliknya.

4. **What is the difference between var and const ?**

	`var` digunakan untuk penyimpanan suatu nilai yang dinamis(berubah-ubah), sedangkan `const` bersifat konstan, atau tidak dapat berubah.

5. **Using the example program as a starting point, write a program that converts
from Fahrenheit into Celsius (C = (F − 32) * 5/9)**

 	```go
 	package main

 	import "fmt"

 	func main() {
 		fmt.Print("Input Fahrenhait: ")

 		var f float64

 		fmt.Scanf("%f", &f)
 		c := (f - 32) * (5.0 / 9.0)

 		fmt.Println(c)
 	}
 	```

6. **Write another program that converts from feet into meters (1 ft = 0.3048 m).**
	```go
	package main

	import "fmt"

	func main() {
		fmt.Print("Input Feet: ")

		var feet float32

		fmt.Scanf("%f", &feet)

		meter := feet * 0.3048

		fmt.Println(meter)
	}
	```


 




