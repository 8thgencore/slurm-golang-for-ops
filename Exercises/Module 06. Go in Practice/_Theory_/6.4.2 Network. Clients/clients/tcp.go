package clients

import (
	"bufio"
	"fmt"
	"net"
)

func MakeTCPConnection() {
	p :=  make([]byte, 2048) //создаем массив байт для чтения ответа от сервера
	conn, err := net.Dial("tcp", "127.0.0.1:4444") //то же самое что и с UDP, но тип соединения TCP
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi TCP Server, How are you doing?") //пишем в conn
	_, err = bufio.NewReader(conn).Read(p) //несмотря на всю разницу форматов, с точки зрения работы клиента в коде - разницы нет
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}