// package main

// import (
// 	"fmt"
// )

// func Slice(a []string, nbrs ...int) []string {
// 	length := len(a)
// 	var start int
// 	var end int

// 	if len(nbrs) == 2 && nbrs[0] < 0 && nbrs[1] < 0 {
// 		start = length + nbrs[0]
// 		end = length + nbrs[1]
// 	} 

// 	if len(nbrs) == 1 && nbrs[0] < 0 {
// 		start = length + nbrs[0]
// 	} 

// 	if len(nbrs) == 2 {
// 		a = a[start:end]
// 	}

// 	if length > end {
// 		a = a[start:end]
// 	} 
// 	return a
// }

// func main() {
// 	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
// 	// fmt.Printf("%#v\n", Slice(a, 1))
// 	fmt.Printf("%#v\n", Slice(a, 2, 4))
// 	// fmt.Printf("%#v\n", Slice(a, -3))
// 	fmt.Printf("%#v\n", Slice(a, -2, -1))
// 	// fmt.Printf("%#v\n", Slice(a, 2, 0))
// }


package main

import (
	"fmt"
)

func Slice(a []string, nbrs... int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	start, end := 0, len(a)
	if len(nbrs) > 0 {
		start = nbrs[0]
		if start < 0 {
			start += len(a)
		}
		if start < 0 || start >= len(a) {
			return nil
		}
	}

	if len(nbrs) > 1 {
		end = nbrs[1]
		if end < 0 {
			end += len(a)
		}
		if end < 0 || end > len(a) {
			end = len(a)
		}
		if start > end {
			return nil
		}
	}

	return a[start:end]
}

func main() {
	a := []string{"coding", "algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 1))             // []string{"algorithm", "ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, 2, 4))          // []string{"ascii", "package"}
	fmt.Printf("%#v\n", Slice(a, -3))            // []string{"ascii", "package", "golang"}
	fmt.Printf("%#v\n", Slice(a, -2, -1))        // []string{"package"}
	fmt.Printf("%#v\n", Slice(a, 2, 0))          // []string(nil)
}
