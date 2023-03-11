package clients

import (
	"bufio"
	"fmt"
	"net"
)

func MakeUDPConnection() {
	p :=  make([]byte, 2048) //создаем массив байт для чтения ответа от сервера
	conn, err := net.Dial("udp", "127.0.0.1:1234") //открываем conn к серверу, используя универсальную команду net.Dial, которой достаточно указать адрес и тип сети
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Hi UDP Server, How are you doing?")//пишем в conn
	_, err = bufio.NewReader(conn).Read(p)//читаем из соединения
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()//обязательно закрываем сокет
}