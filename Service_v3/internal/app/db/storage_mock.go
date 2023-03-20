// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/db/storage.go

// Package db is a generated GoMock package.
package db

import (
	reflect "reflect"
	models "service/internal/app/models"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockStorageInterface is a mock of StorageInterface interface.
type MockStorageInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStorageInterfaceMockRecorder
}

// MockStorageInterfaceMockRecorder is the mock recorder for MockStorageInterface.
type MockStorageInterfaceMockRecorder struct {
	mock *MockStorageInterface
}

// NewMockStorageInterface creates a new mock instance.
func NewMockStorageInterface(ctrl *gomock.Controller) *MockStorageInterface {
	mock := &MockStorageInterface{ctrl: ctrl}
	mock.recorder = &MockStorageInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorageInterface) EXPECT() *MockStorageInterfaceMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockStorageInterface) Add(metric ...models.Metric) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range metric {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Add", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockStorageInterfaceMockRecorder) Add(metric ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockStorageInterface)(nil).Add), metric...)
}

// List mocks base method.
func (m *MockStorageInterface) List(name string, startDate, endDate time.Time, offset, limit int) []models.Metric {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", name, startDate, endDate, offset, limit)
	ret0, _ := ret[0].([]models.Metric)
	return ret0
}

// List indicates an expected call of List.
func (mr *MockStorageInterfaceMockRecorder) List(name, startDate, endDate, offset, limit interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockStorageInterface)(nil).List), name, startDate, endDate, offset, limit)
}