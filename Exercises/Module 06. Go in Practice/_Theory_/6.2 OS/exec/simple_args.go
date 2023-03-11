package exec

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func RunAppWithSimpleArgs() {
	cmd := exec.Command("tr", "a-z", "A-Z") // команда может вызываться со списком аргументов

	cmd.Stdin = strings.NewReader("little slurm goes big") //tr ждет что после запуска появятся какие то вводимые данные, поэтому мы не можем их передать в аргументы

	var out bytes.Buffer
	cmd.Stdout = &out //читаем вывод в любой io.Writer -> в нашем случае Buffer

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("translated phrase: %q\n", out.String()) //выводим результат
}
