package flags

import (
	"flag"
	"fmt"
)

func Simple_flags() {

	wordPtr := flag.String("word", "foo", "a string") // Строковый флаг по ключу word с дефолтным значением foo и описанием a string

	numbPtr := flag.Int("numb", 42, "an int")     // То же самое ток int с именем numb
	boolPtr := flag.Bool("fork", false, "a bool") //Булевский флаг

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var") //Флаги могут быть присвоены уже существующей переменной

	flag.Parse() //без этой команды флаги НЕ поддерживаются вашим приложением, а переменные не будут заполнены. Обязательно вызывать ПОСЛЕ заполнения флагов

	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args()) //остальные переменные не определенные флагами
}
