// Code generated by mockery v2.40.3. DO NOT EDIT.

package ___test

import (
	context "context"

	__ "github.com/gdsc-ys/fluentify-server/gen/proto"

	mock "github.com/stretchr/testify/mock"
)

// MockPronunciationFeedbackServiceServer is an autogenerated mock type for the PronunciationFeedbackServiceServer type
type MockPronunciationFeedbackServiceServer struct {
	mock.Mock
}

type MockPronunciationFeedbackServiceServer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPronunciationFeedbackServiceServer) EXPECT() *MockPronunciationFeedbackServiceServer_Expecter {
	return &MockPronunciationFeedbackServiceServer_Expecter{mock: &_m.Mock}
}

// PronunciationFeedback provides a mock function with given fields: _a0, _a1
func (_m *MockPronunciationFeedbackServiceServer) PronunciationFeedback(_a0 context.Context, _a1 *__.PronunciationFeedbackRequest) (*__.PronunciationFeedbackResponse, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for PronunciationFeedback")
	}

	var r0 *__.PronunciationFeedbackResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *__.PronunciationFeedbackRequest) (*__.PronunciationFeedbackResponse, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *__.PronunciationFeedbackRequest) *__.PronunciationFeedbackResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*__.PronunciationFeedbackResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *__.PronunciationFeedbackRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'PronunciationFeedback'
type MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call struct {
	*mock.Call
}

// PronunciationFeedback is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 *__.PronunciationFeedbackRequest
func (_e *MockPronunciationFeedbackServiceServer_Expecter) PronunciationFeedback(_a0 interface{}, _a1 interface{}) *MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call {
	return &MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call{Call: _e.mock.On("PronunciationFeedback", _a0, _a1)}
}

func (_c *MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call) Run(run func(_a0 context.Context, _a1 *__.PronunciationFeedbackRequest)) *MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*__.PronunciationFeedbackRequest))
	})
	return _c
}

func (_c *MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call) Return(_a0 *__.PronunciationFeedbackResponse, _a1 error) *MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call) RunAndReturn(run func(context.Context, *__.PronunciationFeedbackRequest) (*__.PronunciationFeedbackResponse, error)) *MockPronunciationFeedbackServiceServer_PronunciationFeedback_Call {
	_c.Call.Return(run)
	return _c
}

// mustEmbedUnimplementedPronunciationFeedbackServiceServer provides a mock function with given fields:
func (_m *MockPronunciationFeedbackServiceServer) mustEmbedUnimplementedPronunciationFeedbackServiceServer() {
	_m.Called()
}

// MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'mustEmbedUnimplementedPronunciationFeedbackServiceServer'
type MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call struct {
	*mock.Call
}

// mustEmbedUnimplementedPronunciationFeedbackServiceServer is a helper method to define mock.On call
func (_e *MockPronunciationFeedbackServiceServer_Expecter) mustEmbedUnimplementedPronunciationFeedbackServiceServer() *MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call {
	return &MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call{Call: _e.mock.On("mustEmbedUnimplementedPronunciationFeedbackServiceServer")}
}

func (_c *MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call) Run(run func()) *MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call) Return() *MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call) RunAndReturn(run func()) *MockPronunciationFeedbackServiceServer_mustEmbedUnimplementedPronunciationFeedbackServiceServer_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPronunciationFeedbackServiceServer creates a new instance of MockPronunciationFeedbackServiceServer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPronunciationFeedbackServiceServer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPronunciationFeedbackServiceServer {
	mock := &MockPronunciationFeedbackServiceServer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
