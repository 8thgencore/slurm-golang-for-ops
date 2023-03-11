package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	envLogPath     = "APP_LOGFILE_PATH"
	defaultLogPath = "log.txt"
)

func main() {
	// Создаем обработчик для метода POST /log
	http.HandleFunc("/log", logHandler)

	// Запускаем сервер на порту 8100
	log.Println("Сервер запущен на порту 8100")
	log.Fatal(http.ListenAndServe(":8100", nil))
}

func logHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Метод не поддерживается")
		return
	}

	// Читаем строку из тела запроса
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Ошибка чтения тела запроса")
		return
	}

	defer req.Body.Close()

	log.Println(string(body))

	// Получаем путь к файлу лога из переменной окружения APP_LOGFILE_PATH
	// или используем log.txt по умолчанию
	// os.Setenv(envLogPath, "")
	logfile := os.Getenv(envLogPath)
	if logfile == "" {
		logfile = defaultLogPath
	}

	// Открываем файл лога для дозаписи или создаем новый файл если он не существует
	f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Ошибка открытия файла лога")
		return
	}

	defer f.Close()

	// Записываем строку в файл лога с новой линией
	_, err = f.WriteString(string(body) + "\n")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Ошибка записи в файл лога")
		return
	}

	// Возвращаем код 200 и OK
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
