package main

import "fmt"

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(res)
	return res, length
}

func main() {
	myString := addPrefix("my_string")
	fmt.Println(myString) // prefix_my_string

	myString, err := addPrefixWithErr("error_string")
	fmt.Println(err)      // nil
	fmt.Println(myString) // prefix_error_string

	var f, f2 func(s string) int
	f = func(s string) int { return len(s) }
	f2 = func(s string) int { return 2 }
	//f2 = func() int { return 1 } // ошибка

	_, _ = f(""), f2("")
}
