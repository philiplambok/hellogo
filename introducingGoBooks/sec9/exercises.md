# Tests

1. **Writing a good suite of tests is not always easy, but the process of writing tests often reveals more about a problem than you may at first realize. For example, with our Average function, what happens if you pass in an empty list ( []float64{} )? How could the function be modified to return 0 in this case?**

	Cek jika element kosong, maka return 0 (?)

2. **Write a series of tests for the Min and Max functions you wrote in the previous chapter.**
	
	```go
	package kotori

	import (
		"testing"
	)
	
	func TestMin(t *testing.T) {
		v := Min([]float64{2, 21, 1})

		if v != 1 {
			t.Error("Expected 1, but got ", v)
		}
	}

	func TestMax(t *testing.T) {
		v := Max([]float64{2, 21, 1})

		if v != 21 {
			t.Error("Expected 1, but got ", v)
		}
	}
	```