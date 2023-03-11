package servers

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n") //просто печатаем нашу строку в ResponseWriter, он поддерживает Writer
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {  //получаем хедеры запроса из нашего http.Request
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h) //выводим их построчно
		}
	}
}

func HTTPServer() {
	http.HandleFunc("/hello", hello) //задаем маршруты и обработчик запросов к нашему http ресурсу /hello
	http.HandleFunc("/headers", headers) //задаем маршруты и обработчик запросов к нашему http ресурсу /headers

	err := http.ListenAndServe(":8080", nil) //стартуем сервер, вызовы пакета http напрямую значат что мы используем умолчальный экзамеляр сервиса который http задает за нас
	if err != nil {
		fmt.Println(err)
	}
}