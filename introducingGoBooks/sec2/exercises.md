# Types
1. **How are integers stored on a computer?**

Banyak :)), tergantung bit yang dipakai. Integer di Go menyediakan int8 int16 int32 int64, dan `uint` untuk versi unsigned-nya (bilangan positive saja). untuk pemakaian praktis gunakan `int` saja :)

2. **We know that (in base 10) the largest one-digit number is 9 and the largest two-
digit number is 99. Given that in binary the largest two-digit number is 11 (3),
the largest three-digit number is 111 (7) and the largest four-digit number is 1111
(15), what’s the largest eight-digit number? (Hint: 10 1 −1 = 9 and 10 2 −1 = 99)**

Digit terbesar di bilangan oktal (8)? jawabanya 7. (0 - 7)

3. **Although overpowered for the task, you can use Go as a calculator. Write a pro‐
gram that computes 32,132 × 42,452 and prints it to the terminal (use the * oper‐
ator for multiplication).**
```go
package main

import "fmt"

func main() {
	fmt.Println(32132 * 42452)
}
```

4. **What is a string? How do you find its length?**

Mengunakan fungsi `len("string here..")`

5. **What’s the value of the expression (true && false) || (false && true) || !
(false && false) ?**

- false || false || false
- false (*answer*) 












