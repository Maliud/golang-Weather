package mocks

import mock "github.com/stretchr/testify/mock"


type Logger struct {
	mock.Mock
}

type Logger_Expecter struct {
	mock *mock.Mock
}

func (_m *Logger) EXPECT() *Logger_Expecter {
	return &Logger_Expecter{mock: &_m.Mock}
}

// Debug, verilen alanlarla bir mock fonksiyonu sağlar: msg
func (_m *Logger) Debug(msg string) {
	_m.Called(msg)
}

// Logger_Debug_Call, 'Debug' yöntemi için Run/Return yöntemlerini tür belirterek gölgeleyen bir *mock.Call'dır.

type Logger_Debug_Call struct {
	*mock.Call
}

// Debug, mock.On çağrısını tanımlamak için bir yardımcı yöntemdir
//   - msg string

func(_e *Logger_Expecter) Debug(msg interface{}) *Logger_Debug_Call {
	return &Logger_Debug_Call{Call: _e.mock.On("Debug",msg)}
}

func (_c *Logger_Debug_Call) Run(run func(msg string)) *Logger_Debug_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})

	return _c
}

func (_c *Logger_Debug_Call) Return() *Logger_Debug_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Debug_Call) RunAndReturn(run func(string)) *Logger_Debug_Call {
	_c.Call.Return(run)
	return _c
}

// Error, verilen alanlarla bir mock fonksiyonu sağlar: err


func (_m *Logger) Error(err error) {
	_m.Called(err)
}

// Logger_Error_Call, 'Error' yöntemi için Run/Return yöntemlerini tür belirterek gölgeleyen bir *mock.Call'dır.
type Logger_Error_Call struct {
	*mock.Call
}

// Error, mock.On çağrısını tanımlamak için bir yardımcı yöntemdir
//   - err error

func(_e *Logger_Expecter) Error(err interface{}) *Logger_Error_Call {
	return &Logger_Error_Call{Call: _e.mock.On("Error", err)}
}

func(_c *Logger_Error_Call) Run(run func(err error)) *Logger_Error_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(error))
	})
	return _c
}

func (_c *Logger_Error_Call) Return() *Logger_Error_Call {
	_c.Call.Return()
	return _c
}

func (_c *Logger_Error_Call) RunAndReturn(run func(error)) *Logger_Error_Call {
	_c.Call.Return(run)
	return _c
}


// NewLogger creates a new instance of Logger. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.

func NewLogger(t interface {
	mock.TestingT
	Cleanup(func())
}) *Logger {
	mock := &Logger{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t)})

	return mock
}