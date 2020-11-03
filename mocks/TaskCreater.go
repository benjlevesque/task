// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TaskCreater is an autogenerated mock type for the TaskCreater type
type TaskCreater struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: title
func (_m *TaskCreater) CreateTask(title string) (int, error) {
	ret := _m.Called(title)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(title)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
