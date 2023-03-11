package handlers

import (
	"6_7/example/internals/app/models"
	"6_7/example/internals/app/processors"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

type CarsHandler struct {
	processor *processors.CarsProcessor
}

func NewCarsHandler(processor *processors.CarsProcessor) *CarsHandler { //конструктор
	handler := new(CarsHandler)
	handler.processor = processor
	return handler
}

func (handler *CarsHandler) Create(w http.ResponseWriter, r *http.Request) { // все методы хендлеров имеют одинаковую сигнатуру - две переменных и ничего в ответ
	var newCar models.Car

	err := json.NewDecoder(r.Body).Decode(&newCar) //это обязанность хендлера - готовить переменные в виде понятном внутри приложения
	if err != nil {
		WrapError(w, err) //ошибки должны соотвествовать заданному стандарту, если мы возвращаем JSON - ошибки тоже должны быть JSON (как правило)
		return
	}

	err = handler.processor.CreateCar(newCar) // вызываем метод процессора и отрабатываем ошибки
	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} { //успешный ответ, в общем для сервера формате
		"result" : "OK",
		"data" : "",
	}

	WrapOK(w, m) //здесь возвращаем код 200 и успех
}

func (handler *CarsHandler) List(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	var userIdFilter int64 = 0
	if vars.Get("userid") != "" { //конверсия переменных из GET запросов также отвественность хендлера, если есть какая то простая валидация (тип данных и тд) - также его отвественность
		var err error
		userIdFilter, err = strconv.ParseInt(vars.Get("userid"), 10, 64)
		if err != nil {
			WrapError(w, err)
			return
		}
	}
	list, err := handler.processor.ListCars(userIdFilter, strings.Trim(vars.Get("brand"),"\""), strings.Trim(vars.Get("colour"), "\""), strings.Trim(vars.Get("license_plate"), "\""))
	//для краткости можно собрать структуру с полями по которым мы будем фильтровать

	if err != nil {
		WrapError(w, err)
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : list,
	}

	WrapOK(w, m)
}

func (handler *CarsHandler) Find(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if vars["id"] == "" {
		WrapError(w, errors.New("missing id"))
		return
	}

	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		WrapError(w, err)
		return
	}

	car, err := handler.processor.FindCar(id)

	if err != nil {
		WrapError(w, err)
		return
	}

	var m = map[string]interface{} {
		"result" : "OK",
		"data" : car,
	}

	WrapOK(w, m)
}
