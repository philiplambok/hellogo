# Packages

1. **Why do we use packages?**
	
	Jika sebelumnya kita mengenal function. Fungsi pada golang sudah terbilang sangat powerfull, software akan semakin baik lagi jika mengunakan konsep package, dimana programmer menulis paket-paket secara independent atau mandiri untuk kebutuhan software.

2. **What is the difference between an identifier that starts with a capital letter and one that doesn’t (e.g., Average versus average )?**
	
	Jika menggunakan kata kapital pada fungsi akan membuat fungsi tersebut dapat diakses secara public, sedangkan jika tidak akan terakses secara private.

3. **What is a package alias? How do you make one?**
	
	Walaupun dapat package sudah powefull, namun pasti ada kasus ketika kita mempunyai nama package yang sama, misalnya paket2 tersebut kita abil dari github. Untuk mengatasi masalah tersebut dapat menggunakan alias, `import live "github.com/philiplambok/kotori". Dengan import kita mengubah nama paket aslinya dari kotori menjadi live.

4. **We copied the average function from Chapter 6 to our new package. Create Min and Max functions that find the minimum and maximum values in a slice of float64 s.**
	```go
	package kotori
	// return min element of an array
	func Min(arr []float64) (result float64) {
		result = arr[0]
		for _, val := range arr {
			if val < result {
				result = val
			}
		}

		return
	}

	// return max element of an array
	func Max(arr []float64) (result float64) {
		result = arr[0]
		for _, val := range arr {
			if val > result {
				result = val
			}
		}

		return
	}

	// main folder
	package main

	import "fmt"
	import love "github.com/philiplambok/hellogo/kotori"

	func main() {
		arr := []float64{13, 12}
		fmt.Println(love.Min(arr))
		fmt.Println(love.Max(arr))
	}
	```

5. **How would you document the functions you created in #4?**
	
	Dengan memberikan komentar sebelum function. 
	```go
	// return max element of an array
	func Max(arr []float64) (result float64) {
		result = arr[0]
		for _, val := range arr {
			if val > result {
				result = val
			}
		}

		return
	}
	```
	Untuk melihat dokumentasi, dapat menggunakan perintah go doc `nama_folder_package` `nama_fungsi`
	```
	func Sum(arr []int) (result int)  
		return max element of an array
	``` 



















