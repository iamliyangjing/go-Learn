package main

import "fmt"

func main() {
	// map[key] value
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m) // map[one :1, two:2]
	fmt.Println(len(m))
	fmt.Println(m["one"])
	fmt.Println(m["unknow"])

	r, ok := m["unknow"]
	fmt.Println(r, ok)

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2, m3)
}
