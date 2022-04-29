// Code generated by mockery v2.10.4. DO NOT EDIT.

package objazure

import (
	azblob "github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	mock "github.com/stretchr/testify/mock"
)

// mockContainerAPI is an autogenerated mock type for the containerAPI type
type mockContainerAPI struct {
	mock.Mock
}

// GetListBlobsFlatPagerAPI provides a mock function with given fields: options
func (_m *mockContainerAPI) GetListBlobsFlatPagerAPI(options azblob.ContainerListBlobsFlatOptions) listBlobsPagerAPI {
	ret := _m.Called(options)

	var r0 listBlobsPagerAPI
	if rf, ok := ret.Get(0).(func(azblob.ContainerListBlobsFlatOptions) listBlobsPagerAPI); ok {
		r0 = rf(options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listBlobsPagerAPI)
		}
	}

	return r0
}

// GetListBlobsHierarchyPagerAPI provides a mock function with given fields: delimiter, options
func (_m *mockContainerAPI) GetListBlobsHierarchyPagerAPI(delimiter string, options azblob.ContainerListBlobsHierarchyOptions) listBlobsPagerAPI {
	ret := _m.Called(delimiter, options)

	var r0 listBlobsPagerAPI
	if rf, ok := ret.Get(0).(func(string, azblob.ContainerListBlobsHierarchyOptions) listBlobsPagerAPI); ok {
		r0 = rf(delimiter, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listBlobsPagerAPI)
		}
	}

	return r0
}

// ToBlobAPI provides a mock function with given fields: blob
func (_m *mockContainerAPI) ToBlobAPI(blob string) (blobAPI, error) {
	ret := _m.Called(blob)

	var r0 blobAPI
	if rf, ok := ret.Get(0).(func(string) blobAPI); ok {
		r0 = rf(blob)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(blobAPI)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(blob)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
