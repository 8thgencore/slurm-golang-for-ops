package envs

import (
	"fmt"
	"os"
	"strings"
)

func Simple_envs() {

	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO")) //получаем переменные окружения
	fmt.Println("BAR:", os.Getenv("BAR")) //здесь дефолтное значение

	fmt.Println()
	for _, e := range os.Environ() { //выводим все что видит приложение в парах
		pair := strings.SplitN(e, "=", 2) // переменные окружения предоставляются системой в виде ИМЯ=значение, поэтому их можно разделить по =
		fmt.Println(pair[0])
		fmt.Println(pair[1])
	}
}
