// Code generated by mockery v2.40.3. DO NOT EDIT.

package ___test

import mock "github.com/stretchr/testify/mock"

// MockisFeedbackHistoryDTO_Feedback is an autogenerated mock type for the isFeedbackHistoryDTO_Feedback type
type MockisFeedbackHistoryDTO_Feedback struct {
	mock.Mock
}

type MockisFeedbackHistoryDTO_Feedback_Expecter struct {
	mock *mock.Mock
}

func (_m *MockisFeedbackHistoryDTO_Feedback) EXPECT() *MockisFeedbackHistoryDTO_Feedback_Expecter {
	return &MockisFeedbackHistoryDTO_Feedback_Expecter{mock: &_m.Mock}
}

// isFeedbackHistoryDTO_Feedback provides a mock function with given fields:
func (_m *MockisFeedbackHistoryDTO_Feedback) isFeedbackHistoryDTO_Feedback() {
	_m.Called()
}

// MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'isFeedbackHistoryDTO_Feedback'
type MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call struct {
	*mock.Call
}

// isFeedbackHistoryDTO_Feedback is a helper method to define mock.On call
func (_e *MockisFeedbackHistoryDTO_Feedback_Expecter) isFeedbackHistoryDTO_Feedback() *MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call {
	return &MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call{Call: _e.mock.On("isFeedbackHistoryDTO_Feedback")}
}

func (_c *MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call) Run(run func()) *MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call) Return() *MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call) RunAndReturn(run func()) *MockisFeedbackHistoryDTO_Feedback_isFeedbackHistoryDTO_Feedback_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockisFeedbackHistoryDTO_Feedback creates a new instance of MockisFeedbackHistoryDTO_Feedback. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockisFeedbackHistoryDTO_Feedback(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockisFeedbackHistoryDTO_Feedback {
	mock := &MockisFeedbackHistoryDTO_Feedback{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}