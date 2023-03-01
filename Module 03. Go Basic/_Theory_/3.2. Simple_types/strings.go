package main

import "fmt"

func main()  {
	var myString string // ""

	var hello string = "\tHello\n"
	var title string = `\tHello\n`

	word := "my string" // автоматическое определение типа


	str := "你好"
	str = "Привет"

	var b byte = 'c' // равно 99
	var r rune = '你' // равно 20320

	str = "12345"
	fmt.Println(str[0]) // 49
	fmt.Println(str[1]) // 50
	_ = len(str) // 5

	str = "строка" // 6 символов
	fmt.Println(len(str)) // 12


	//str[from:to]
	newStr := str[2:4] // "34"
	newStr = str[:4] // "1234"
	newStr = str[2:] // "345"
	// from - 0, to - len(str)

	word = "Hello"
	hello = word + " world!" // "Hello world"
	word[0] = "h" // ошибка


	isBigger := "aba" > "abc"
	isBigger = "aba" < "abc"
	isEqual := "aba" == "abc"



	fmt.Println(title, myString, hello, str, b, r, word, newStr)
}
