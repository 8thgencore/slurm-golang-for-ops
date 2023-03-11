package servers
import (
	"bufio"
	"fmt"
	"net"
	"strings"
)


func sendTCPResponse(conn net.Conn) {
	netData, err := bufio.NewReader(conn).ReadString('\n') //читаем сообщение из Conn
	if err != nil {
		fmt.Println(err)
		return
	}

	temp := strings.TrimSpace(netData) // убираем лишние пробелы
	fmt.Println(temp)

	_,err = conn.Write([]byte("From server: Hello I got your message ")) //пишем в наш Conn, больше ничего не надо
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}


func TCPServer() {
	addr := net.TCPAddr{ //Создаем адрес практически аналогично UDP, но с использованием структуры TCPAddr
		Port: 1234,
		IP: net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenTCP("tcp", &addr) //слушаем входные TCP соединения

	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	defer ser.Close() //закрываем соединение
	for {
		conn, err := ser.Accept() //поскольку TCP создает соединение, мы должны его принять, и только после этого получаем обьект Conn для чтения и записи
		if err != nil {
			fmt.Println(err)
			return
		}
		go sendTCPResponse(conn) //пишем в наш Conn, адрес уже не нужен, поскольку conn знает про клиента
	}
}