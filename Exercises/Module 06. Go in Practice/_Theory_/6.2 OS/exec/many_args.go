package exec

import (
	"fmt"
	"os/exec"
)

func RunMultipleArgs() {

	prg := "echo"

	arg1 := "there"
	arg2 := "are slurms"
	arg3 := "in Slurmland"

	cmd := exec.Command(prg, arg1, arg2, arg3) // аргументы последовательно выводятся для прогрммы, в случае echo это просто одна строка
	stdout, err := cmd.Output() // читаем возврат команды одной строкой

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}
