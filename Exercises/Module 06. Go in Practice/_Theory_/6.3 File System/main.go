package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	write("test.txt")
}

func write(filepath string) {
	f, err := os.Create(filepath) //создаем файл и получаем на него ссылку
	check(err)

	defer f.Close() //закрываем файл в конце, после чтения или записи, обязательно делаем это для корректной работы с нашей ОС

	d2 := []byte{115, 111, 109, 101, 10} //пишем байты напрямую
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n") //пишем строку
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	_ = f.Sync() //при выполнении команды Sync содержимое файла запишется на диск, до этого содержимое файла может находиться в кэше системы и данные потеряются, при, скажем, внезапной перезагрузке

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n") //пишем в буфер
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() //сохраняем содержимое буфера
}