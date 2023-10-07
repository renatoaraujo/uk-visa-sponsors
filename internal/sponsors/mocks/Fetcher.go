// Code generated by mockery v2.34.1. DO NOT EDIT.

package mocks_test

import mock "github.com/stretchr/testify/mock"

// Fetcher is an autogenerated mock type for the Fetcher type
type Fetcher struct {
	mock.Mock
}

type Fetcher_Expecter struct {
	mock *mock.Mock
}

func (_m *Fetcher) EXPECT() *Fetcher_Expecter {
	return &Fetcher_Expecter{mock: &_m.Mock}
}

// FetchData provides a mock function with given fields: url
func (_m *Fetcher) FetchData(url string) ([]byte, error) {
	ret := _m.Called(url)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(url)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(url)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(url)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetcher_FetchData_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchData'
type Fetcher_FetchData_Call struct {
	*mock.Call
}

// FetchData is a helper method to define mock.On call
//   - url string
func (_e *Fetcher_Expecter) FetchData(url interface{}) *Fetcher_FetchData_Call {
	return &Fetcher_FetchData_Call{Call: _e.mock.On("FetchData", url)}
}

func (_c *Fetcher_FetchData_Call) Run(run func(url string)) *Fetcher_FetchData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Fetcher_FetchData_Call) Return(_a0 []byte, _a1 error) *Fetcher_FetchData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Fetcher_FetchData_Call) RunAndReturn(run func(string) ([]byte, error)) *Fetcher_FetchData_Call {
	_c.Call.Return(run)
	return _c
}

// FetchDataSource provides a mock function with given fields:
func (_m *Fetcher) FetchDataSource() (string, error) {
	ret := _m.Called()

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func() (string, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fetcher_FetchDataSource_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchDataSource'
type Fetcher_FetchDataSource_Call struct {
	*mock.Call
}

// FetchDataSource is a helper method to define mock.On call
func (_e *Fetcher_Expecter) FetchDataSource() *Fetcher_FetchDataSource_Call {
	return &Fetcher_FetchDataSource_Call{Call: _e.mock.On("FetchDataSource")}
}

func (_c *Fetcher_FetchDataSource_Call) Run(run func()) *Fetcher_FetchDataSource_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Fetcher_FetchDataSource_Call) Return(_a0 string, _a1 error) *Fetcher_FetchDataSource_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Fetcher_FetchDataSource_Call) RunAndReturn(run func() (string, error)) *Fetcher_FetchDataSource_Call {
	_c.Call.Return(run)
	return _c
}

// NewFetcher creates a new instance of Fetcher. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFetcher(t interface {
	mock.TestingT
	Cleanup(func())
}) *Fetcher {
	mock := &Fetcher{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}