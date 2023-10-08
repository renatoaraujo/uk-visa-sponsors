// Code generated by mockery v2.34.1. DO NOT EDIT.

package mocks_test

import mock "github.com/stretchr/testify/mock"

// HttpUtils is an autogenerated mock type for the HttpUtils type
type HttpUtils struct {
	mock.Mock
}

type HttpUtils_Expecter struct {
	mock *mock.Mock
}

func (_m *HttpUtils) EXPECT() *HttpUtils_Expecter {
	return &HttpUtils_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: resourcePath
func (_m *HttpUtils) Get(resourcePath string) ([]byte, error) {
	ret := _m.Called(resourcePath)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(resourcePath)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(resourcePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(resourcePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HttpUtils_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type HttpUtils_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - resourcePath string
func (_e *HttpUtils_Expecter) Get(resourcePath interface{}) *HttpUtils_Get_Call {
	return &HttpUtils_Get_Call{Call: _e.mock.On("Get", resourcePath)}
}

func (_c *HttpUtils_Get_Call) Run(run func(resourcePath string)) *HttpUtils_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *HttpUtils_Get_Call) Return(_a0 []byte, _a1 error) *HttpUtils_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *HttpUtils_Get_Call) RunAndReturn(run func(string) ([]byte, error)) *HttpUtils_Get_Call {
	_c.Call.Return(run)
	return _c
}

// NewHttpUtils creates a new instance of HttpUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHttpUtils(t interface {
	mock.TestingT
	Cleanup(func())
}) *HttpUtils {
	mock := &HttpUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
