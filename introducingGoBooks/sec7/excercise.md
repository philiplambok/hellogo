# Structs, Methods and Interface

1. **What’s the difference between a method and a function?**

	Function adalah fungsi biasa yang dipakai secara modular terhadap sebuah paket(package), sedangkan method adalah fungsi yang menempel(embedded) terhapat suatu struck. 

2. **Why would you use an embedded anonymous field instead of a normal named
field?**
	
	Agar dapat menggunakan struct (Relasi) lebih powerfull. Dengan memberikan anony field anda melakukan store metode terhadap objek yang di store.

3. **Add a new perimeter method to the Shape interface to calculate the perimeter of
a shape. Implement the method for Circle and Rectangle .**
	```go
	package main
	import "fmt"

	type Shape interface {
		perim() float64
	}
	
	type Circle struct {
		diameter float64
	}

	type Rectangle struct {
		base, height float64
	}

	func (r Rectangle) perim() float64 {
		return 2 * r.base + 2 * r.height
	}

	func (c Circle) perim() float64 {
		return 2 * 3.14 * r.base
	}

	func totalPerim(shapes ...Shape) float64 {
		total := 0.0
		
		for _, shape := range Shape {
			total += shape.perim()
		}

		return total
	}

	func main() {
		c := Circle{7}
		r := Rectangle{4,2}

		fmt.Println(`Perim Circle `, c.perim())
		fmt.Println(`Perim Rectangle`, r.perim())
		fmt.Println(`Total Perim`, r.totalPerim())
	}

	```













