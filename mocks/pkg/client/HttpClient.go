package mocks

import mock "github.com/stretchr/testify/mock"

type HttpClient struct {
	mock.Mock
}

type HttpClient_Expecter struct {
	mock *mock.Mock
}

func (_m *HttpClient) EXPECT() *HttpClient_Expecter{
	return &HttpClient_Expecter{mock: &_m.Mock}
}

// GetWeatherData Verilen alanlarla bir mock işlevi sağlar: boylam, enlem

func(_m *HttpClient) GetWeatherData(longitude float64, latitude float64) (string, error) {
	ret := _m.Called(longitude, latitude)

	if len(ret) == 0 {
		panic("GetWeatherData için dönüş değeri belirtilmemiş")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(float64, float64) (string, error)); ok {
		return rf(longitude, latitude)
	}

	if rf, ok := ret.Get(0).(func(float64, float64) string); ok {
		 r0 = rf(longitude, latitude)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(float64, float64) error); ok {
		r1 = rf(longitude, latitude)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HttpClient_GetWeatherData_Call, 'GetWeatherData' yöntemi için Run/Return yöntemlerini tür belirterek gölgeleyen bir *mock.Call'dır.

type HttpClient_GetWeatherData_Call struct {
	*mock.Call
}

// GetWeatherData, mock.On çağrısını tanımlamak için bir yardımcı yöntemdir
//   - boylam float64
//   - enlem float64


func (_e *HttpClient_Expecter) GetWeatherData(longitude interface{}, latitude interface{}) *HttpClient_GetWeatherData_Call {
	return &HttpClient_GetWeatherData_Call{Call: _e.mock.On("GetWeatherData", longitude, latitude)}
}

func (_c *HttpClient_GetWeatherData_Call) Run(run func(longitude float64, latitude float64)) *HttpClient_GetWeatherData_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64), args[1].(float64))
	})

	return _c
}

func (_c *HttpClient_GetWeatherData_Call) Return(_a0 string, _a1 error) *HttpClient_GetWeatherData_Call {
	_c.Call.Return(_a0, _a1)
	return _c
} 

func (_c *HttpClient_GetWeatherData_Call) RunAndReturn(run func(float64, float64) (string, error)) *HttpClient_GetWeatherData_Call {
	_c.Call.Return(run)
	return _c
}

// NewHttpClient, yeni bir HttpClient örneği oluşturur. Ayrıca, mock üzerinde bir test arayüzü ve mock beklentilerini doğrulamak için bir temizleme işlevi kaydeder.
// İlk argüman genellikle *testing.T değeridir.


func NewHttpClient(t interface{
	mock.TestingT
	Cleanup(func())
}) *HttpClient {
	mock := &HttpClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t)})
	return mock
}