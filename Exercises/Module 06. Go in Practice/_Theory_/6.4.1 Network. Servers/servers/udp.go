package servers
import (
	"fmt"
	"net"
)


func sendResponse(conn *net.UDPConn, addr *net.UDPAddr) {
	_,err := conn.WriteToUDP([]byte("From server: Hello I got your message "), addr) //просто отправляем ответ обратно, записывая его в канал
	if err != nil {
		fmt.Printf("Couldn't send response %v", err)
	}
}


func UDPServer() {
	p := make([]byte, 2048) // создаем буфер для чтения входных данных от нашего клиента
	addr := net.UDPAddr{ //создаем UDP адрес для запуска сервера
		Port: 1234, // порт
		IP: net.ParseIP("127.0.0.1"), // ParseIP разберет адрес из строки в набор байтов для использования системой
	}
	ser, err := net.ListenUDP("udp", &addr) //открываем соединение по адресу и получаем UDPConn для чтения данных
	defer ser.Close() // закрываем соединение
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}
	for {
		_,remoteaddr,err := ser.ReadFromUDP(p) //ждем данных  из сокета, когда они поступили и читаем в p, функция вернет n - количество байт, адрес вызвающего для ответа, и возможную ошибку
		fmt.Printf("Read a message from %v %s \n", remoteaddr, p)
		if err !=  nil {
			fmt.Printf("Some error  %v", err)
			continue
		}
		go sendResponse(ser, remoteaddr) //отвечаем в отдельной горутине
	}
}