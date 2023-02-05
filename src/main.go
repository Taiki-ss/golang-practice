package main

import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(writer, "Hello World")
}

func main() {
	// http.HandleFunc("/", handler)
	// http.ListenAndServe(":8080", nil)

	f := 1.11
	fmt.Printf("%d\n", int(f))

	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v\n", m, m)

	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}

	var minNumber int
	for i := 0; i < len(l); i++ {
		if i == 0 {
			minNumber = l[i]
			continue
		}

		if minNumber > l[i] {
			minNumber = l[i]
		}
	}

	fmt.Println(minNumber)
}
