package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
)


func RequestLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Request: %s Method: %s",r.RequestURI, r.Method)
			next.ServeHTTP(w, r) //главная функция для продолжения работы, без нее ваш Middleware поломает ответы на запросы, поскольку не передаст управление функциям из Router
	}) // задача Handler - обработка запросов, поэтому Middleware должен вернуть handler, мы используем HandlerFunc для простоты
}