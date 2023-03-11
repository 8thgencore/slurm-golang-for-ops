package exec

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
)

func upper(data string) string {
	return strings.ToUpper(data)
}

func OutputPipe() {
	cmd := exec.Command("echo", "piping slurms")

	stdout, err := cmd.StdoutPipe() //открываем пайп для продолжительного чтения вывода

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(stdout) // читаем все из пайпа

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil { // Wait закроет пайп по сигналу завершения вывода, однако подождет пока все чтение не завершится
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(data))
}
