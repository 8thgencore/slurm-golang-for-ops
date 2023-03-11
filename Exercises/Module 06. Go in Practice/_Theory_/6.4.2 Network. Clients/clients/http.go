package clients

import (
	"bufio"
	"fmt"
	"net/http"
)

func HTTPClient() {
	resp, err := http.Get("http://localhost:8080/hello") //тут уже не получится использовать dial, поскольку нам нужно работать на другом уровне, используем GET для получения ресурса /hello
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()//поскольку http работает через tcp, нам нужно закрыть соединение, это делается через метод Body

	fmt.Println("Response status:", resp.Status)//узнаем статус соединения, должно быть 200 или 201

	scanner := bufio.NewScanner(resp.Body) //Body - это интерфейс Reader, его можно использовать для чтения данных построчно
	for i := 0; scanner.Scan() && i < 5; i++ { //читаем данные построчно, ограничиваем чтение пятью строками (в реальных http документах может быть намного больше, можно убрать i < 5 для чтения всего документа)
		fmt.Println(scanner.Text())//выводим текст, сканер будет читать строку за строкой, следовательно выводиться у нас будет по одной строке
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
