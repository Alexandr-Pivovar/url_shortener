// Code generated by mockery v1.0.0. DO NOT EDIT.

package app

import mock "github.com/stretchr/testify/mock"

// MockUrlRepository is an autogenerated mock type for the UrlRepository type
type MockUrlRepository struct {
	mock.Mock
}

// Get provides a mock function with given fields: key
func (_m *MockUrlRepository) Get(key string) (string, error) {
	ret := _m.Called(key)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: key, url
func (_m *MockUrlRepository) Save(key string, url string) error {
	ret := _m.Called(key, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(key, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
