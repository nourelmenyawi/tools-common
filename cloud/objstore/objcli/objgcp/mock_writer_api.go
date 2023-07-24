// Code generated by mockery v2.32.0. DO NOT EDIT.

package objgcp

import mock "github.com/stretchr/testify/mock"

// mockWriterAPI is an autogenerated mock type for the writerAPI type
type mockWriterAPI struct {
	mock.Mock
}

// Close provides a mock function with given fields:
func (_m *mockWriterAPI) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SendCRC provides a mock function with given fields: crc
func (_m *mockWriterAPI) SendCRC(crc uint32) {
	_m.Called(crc)
}

// SendMD5 provides a mock function with given fields: md5
func (_m *mockWriterAPI) SendMD5(md5 []byte) {
	_m.Called(md5)
}

// Write provides a mock function with given fields: p
func (_m *mockWriterAPI) Write(p []byte) (int, error) {
	ret := _m.Called(p)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte) (int, error)); ok {
		return rf(p)
	}
	if rf, ok := ret.Get(0).(func([]byte) int); ok {
		r0 = rf(p)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// newMockWriterAPI creates a new instance of mockWriterAPI. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockWriterAPI(t interface {
	mock.TestingT
	Cleanup(func())
},
) *mockWriterAPI {
	mock := &mockWriterAPI{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
