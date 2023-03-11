package exec

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Piping() {
	c1 := exec.Command("cat", "test.txt")
	c2 := exec.Command("grep", "Reading")
	//определяем нашу последовательность команд

	r, w := io.Pipe() //создаем пайп на запись и чтение
	c1.Stdout = w // одна команда должна в пайп писать
	c2.Stdin = r //вторая из него читать

	var b2 bytes.Buffer //собираем вывод второй команд
	c2.Stdout = &b2

	c1.Start()
	c2.Start()
	c1.Wait() //надо чтобы первая завершилась раньше второй
	w.Close() // закрываем вывод
	c2.Wait() // завершаем вторую
	io.Copy(os.Stdout, &b2) //копируем вывод, поскольку Stdout это обычный Writer, его можно скопировать в буфер

	fmt.Println(b2)
}