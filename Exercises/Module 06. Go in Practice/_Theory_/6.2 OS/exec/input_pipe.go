package exec

import (
	"fmt"
	"io"
	"log"
	"os/exec"
)

func InputPiping() {

	cmd := exec.Command("cat")
	stdin, err := cmd.StdinPipe() //открываем ввод команды (stdin pipe), в который можем писать откуда угодно
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "an old slurm") // пишем в stdin из горутины
	}()

	out, err := cmd.CombinedOutput() // запускаем и возвращаем вывод
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", out)
}
