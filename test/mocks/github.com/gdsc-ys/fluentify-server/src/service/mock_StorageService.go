// Code generated by mockery v2.40.3. DO NOT EDIT.

package service_test

import mock "github.com/stretchr/testify/mock"

// MockStorageService is an autogenerated mock type for the StorageService type
type MockStorageService struct {
	mock.Mock
}

type MockStorageService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStorageService) EXPECT() *MockStorageService_Expecter {
	return &MockStorageService_Expecter{mock: &_m.Mock}
}

// GetFile provides a mock function with given fields: filePath
func (_m *MockStorageService) GetFile(filePath string) ([]byte, error) {
	ret := _m.Called(filePath)

	if len(ret) == 0 {
		panic("no return value specified for GetFile")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]byte, error)); ok {
		return rf(filePath)
	}
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(filePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageService_GetFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFile'
type MockStorageService_GetFile_Call struct {
	*mock.Call
}

// GetFile is a helper method to define mock.On call
//   - filePath string
func (_e *MockStorageService_Expecter) GetFile(filePath interface{}) *MockStorageService_GetFile_Call {
	return &MockStorageService_GetFile_Call{Call: _e.mock.On("GetFile", filePath)}
}

func (_c *MockStorageService_GetFile_Call) Run(run func(filePath string)) *MockStorageService_GetFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStorageService_GetFile_Call) Return(_a0 []byte, _a1 error) *MockStorageService_GetFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageService_GetFile_Call) RunAndReturn(run func(string) ([]byte, error)) *MockStorageService_GetFile_Call {
	_c.Call.Return(run)
	return _c
}

// GetFileUrl provides a mock function with given fields: filePath
func (_m *MockStorageService) GetFileUrl(filePath string) (string, error) {
	ret := _m.Called(filePath)

	if len(ret) == 0 {
		panic("no return value specified for GetFileUrl")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (string, error)); ok {
		return rf(filePath)
	}
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(filePath)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(filePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageService_GetFileUrl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetFileUrl'
type MockStorageService_GetFileUrl_Call struct {
	*mock.Call
}

// GetFileUrl is a helper method to define mock.On call
//   - filePath string
func (_e *MockStorageService_Expecter) GetFileUrl(filePath interface{}) *MockStorageService_GetFileUrl_Call {
	return &MockStorageService_GetFileUrl_Call{Call: _e.mock.On("GetFileUrl", filePath)}
}

func (_c *MockStorageService_GetFileUrl_Call) Run(run func(filePath string)) *MockStorageService_GetFileUrl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockStorageService_GetFileUrl_Call) Return(_a0 string, _a1 error) *MockStorageService_GetFileUrl_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageService_GetFileUrl_Call) RunAndReturn(run func(string) (string, error)) *MockStorageService_GetFileUrl_Call {
	_c.Call.Return(run)
	return _c
}

// UploadFile provides a mock function with given fields: file, userId
func (_m *MockStorageService) UploadFile(file []byte, userId string) (string, error) {
	ret := _m.Called(file, userId)

	if len(ret) == 0 {
		panic("no return value specified for UploadFile")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, string) (string, error)); ok {
		return rf(file, userId)
	}
	if rf, ok := ret.Get(0).(func([]byte, string) string); ok {
		r0 = rf(file, userId)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func([]byte, string) error); ok {
		r1 = rf(file, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStorageService_UploadFile_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UploadFile'
type MockStorageService_UploadFile_Call struct {
	*mock.Call
}

// UploadFile is a helper method to define mock.On call
//   - file []byte
//   - userId string
func (_e *MockStorageService_Expecter) UploadFile(file interface{}, userId interface{}) *MockStorageService_UploadFile_Call {
	return &MockStorageService_UploadFile_Call{Call: _e.mock.On("UploadFile", file, userId)}
}

func (_c *MockStorageService_UploadFile_Call) Run(run func(file []byte, userId string)) *MockStorageService_UploadFile_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].([]byte), args[1].(string))
	})
	return _c
}

func (_c *MockStorageService_UploadFile_Call) Return(_a0 string, _a1 error) *MockStorageService_UploadFile_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockStorageService_UploadFile_Call) RunAndReturn(run func([]byte, string) (string, error)) *MockStorageService_UploadFile_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockStorageService creates a new instance of MockStorageService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockStorageService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockStorageService {
	mock := &MockStorageService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}