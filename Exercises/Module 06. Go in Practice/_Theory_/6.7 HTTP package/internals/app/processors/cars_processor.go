package processors

import (
	"6_7/example/internals/app/db"
	"6_7/example/internals/app/models"
	"errors"
)

type CarsProcessor struct {
	storage *db.CarsStorage
}

func NewCarsProcessor(storage *db.CarsStorage) *CarsProcessor {
	processor := new(CarsProcessor)
	processor.storage = storage
	return processor
}

func (processor *CarsProcessor) CreateCar(car models.Car) error { //Процессор выполняет внутреннюю бизнес логику и валидирует переменные в соотвествии с ней
	if car.Colour == "" {
		return errors.New("colour should not be empty")
	} // нельзя создать машину без цвета

	if car.Brand == "" {
		return errors.New("brand should not be empty")
	} //обязательно должен быть указан бренд

	if car.Owner.Id < 0 {
		return errors.New("user shall be filled")
	} //машин без владельца на свете тоже не бывает

	return processor.storage.CreateCar(car)
}

func (processor *CarsProcessor) FindCar(id int64) (models.Car, error) {
	car := processor.storage.GetCarById(id)

	if car.Id != id { //опять же внутренняя бизнес логика
		return car, errors.New("car not found")
	}

	return car, nil
}

func (processor *CarsProcessor) ListCars(userId int64, brandFilter string, colourFilter string, licenseFilter string) ([]models.Car, error) {
	return processor.storage.GetCarsList(userId, brandFilter, colourFilter, licenseFilter), nil
}