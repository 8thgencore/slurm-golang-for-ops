package processors

import (
	"service/internal/app/db"
	"service/internal/app/models"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestAdd(t *testing.T) {
	ctrl := gomock.NewController(t) // Создаем контроллер
	defer ctrl.Finish()             // Закрываем контроллер в конце

	// Создаем фиктивное хранилище типа db.MockStorageInterface
	mockStorage := db.NewMockStorageInterface(ctrl)

	// Создаем экземпляр MetricsProcessor с фиктивным хранилищем
	processor := NewMetricsProcessor(mockStorage)

	// Создаем тестовые данные
	testCases := []struct {
		metric models.Metric // Входная метрика
		err    error         // Ожидаемая ошибка
	}{
		{
			metric: models.Metric{
				Name:  "temperature",
				Value: "25",
				Date:  time.Now(),
			},
			err: nil,
		},
		{
			metric: models.Metric{
				Name:  "",
				Value: "25",
				Date:  time.Now(),
			},
			err: ErrEmptyName,
		},
		{
			metric: models.Metric{
				Name:  "temperature",
				Value: "",
				Date:  time.Now(),
			},
			err: ErrEmptyValue,
		},
		{
			metric: models.Metric{
				Name:  "temperature",
				Value: "25",
				Date:  time.Time{},
			},
			err: ErrEmptyDate,
		},
		// Добавить другие случаи по необходимости
	}

	// Перебираем тестовые данные
	for _, tc := range testCases {
		if tc.err == nil { // Если ожидается успешный вызов метода Add
			mockStorage.EXPECT().Add(tc.metric).Return(nil) // Настраиваем ожидание вызова Add с входной метрикой и возвратом nil ошибки
		}
		err := processor.Add(tc.metric) // Вызываем метод Add с входной метрикой
		if err != tc.err {              // Сравниваем полученную ошибку с ожидаемой
			t.Errorf("Expected error %v, got %v", tc.err, err) // Выводим сообщение об ошибке, если они не совпадают
		}
	}
}
