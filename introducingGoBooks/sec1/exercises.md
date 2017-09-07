# Getting Started

1. **What is whitespace?**

	Whitespace adalah spasi kosong, atau enter kosong yang digunakan saat pembuatan program (Development) untuk membuat kode mudah dibaca.

2. **What is a comment? What are the two ways of writing a comment?**

	Comment berguna untuk dokumentasi program. Ada dua cara untuk menulis komentar, pertama dengan `// Comment Here` untuk satu baris komentar dan `/* */` Untuk komentar yang lebih dari satu baris.

3. **Our program began with package main . What would the files in the fmt package
begin with ?**

	Untuk dapat memakai paket fmt(format) harus melalukan import terlabih dahulu `import "fmt"`. Dan file-file dalam pake fmt akan dimulai dengan pendeklarasian paketnya terlebih dahulu `package fmt`.

4. **We used the Println function defined in the fmt package. If you wanted to use
the Exit function from the os package, what would you need to do?**

	Menggunakan perintah import pada pake main `import "os"`. Lalu dapat menggunakan fungsi main dengan `os.Exit()`.

5. **Modify the program we wrote so that instead of printing Hello, World it prints
Hello, my name is followed by your name.**

	```go
	package main

	import "fmt"

	func main(){
		var name string = "Philip lambok"
		fmt.Println(name)
	}
	```