// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// TaskToggler is an autogenerated mock type for the TaskToggler type
type TaskToggler struct {
	mock.Mock
}

// ToggleTask provides a mock function with given fields: _a0, _a1
func (_m *TaskToggler) ToggleTask(_a0 int, _a1 bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
