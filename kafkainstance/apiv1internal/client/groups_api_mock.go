// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package kafkainstanceclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that GroupsApiMock does implement GroupsApi.
// If this is not the case, regenerate this file with moq.
var _ GroupsApi = &GroupsApiMock{}

// GroupsApiMock is a mock implementation of GroupsApi.
//
// 	func TestSomethingThatUsesGroupsApi(t *testing.T) {
//
// 		// make and configure a mocked GroupsApi
// 		mockedGroupsApi := &GroupsApiMock{
// 			DeleteConsumerGroupByIdFunc: func(ctx _context.Context, consumerGroupId string) ApiDeleteConsumerGroupByIdRequest {
// 				panic("mock out the DeleteConsumerGroupById method")
// 			},
// 			DeleteConsumerGroupByIdExecuteFunc: func(r ApiDeleteConsumerGroupByIdRequest) (*_nethttp.Response, error) {
// 				panic("mock out the DeleteConsumerGroupByIdExecute method")
// 			},
// 			GetConsumerGroupByIdFunc: func(ctx _context.Context, consumerGroupId string) ApiGetConsumerGroupByIdRequest {
// 				panic("mock out the GetConsumerGroupById method")
// 			},
// 			GetConsumerGroupByIdExecuteFunc: func(r ApiGetConsumerGroupByIdRequest) (ConsumerGroup, *_nethttp.Response, error) {
// 				panic("mock out the GetConsumerGroupByIdExecute method")
// 			},
// 			GetConsumerGroupsFunc: func(ctx _context.Context) ApiGetConsumerGroupsRequest {
// 				panic("mock out the GetConsumerGroups method")
// 			},
// 			GetConsumerGroupsExecuteFunc: func(r ApiGetConsumerGroupsRequest) (ConsumerGroupList, *_nethttp.Response, error) {
// 				panic("mock out the GetConsumerGroupsExecute method")
// 			},
// 			ResetConsumerGroupOffsetFunc: func(ctx _context.Context, consumerGroupId string) ApiResetConsumerGroupOffsetRequest {
// 				panic("mock out the ResetConsumerGroupOffset method")
// 			},
// 			ResetConsumerGroupOffsetExecuteFunc: func(r ApiResetConsumerGroupOffsetRequest) (ConsumerGroupResetOffsetResult, *_nethttp.Response, error) {
// 				panic("mock out the ResetConsumerGroupOffsetExecute method")
// 			},
// 		}
//
// 		// use mockedGroupsApi in code that requires GroupsApi
// 		// and then make assertions.
//
// 	}
type GroupsApiMock struct {
	// DeleteConsumerGroupByIdFunc mocks the DeleteConsumerGroupById method.
	DeleteConsumerGroupByIdFunc func(ctx _context.Context, consumerGroupId string) ApiDeleteConsumerGroupByIdRequest

	// DeleteConsumerGroupByIdExecuteFunc mocks the DeleteConsumerGroupByIdExecute method.
	DeleteConsumerGroupByIdExecuteFunc func(r ApiDeleteConsumerGroupByIdRequest) (*_nethttp.Response, error)

	// GetConsumerGroupByIdFunc mocks the GetConsumerGroupById method.
	GetConsumerGroupByIdFunc func(ctx _context.Context, consumerGroupId string) ApiGetConsumerGroupByIdRequest

	// GetConsumerGroupByIdExecuteFunc mocks the GetConsumerGroupByIdExecute method.
	GetConsumerGroupByIdExecuteFunc func(r ApiGetConsumerGroupByIdRequest) (ConsumerGroup, *_nethttp.Response, error)

	// GetConsumerGroupsFunc mocks the GetConsumerGroups method.
	GetConsumerGroupsFunc func(ctx _context.Context) ApiGetConsumerGroupsRequest

	// GetConsumerGroupsExecuteFunc mocks the GetConsumerGroupsExecute method.
	GetConsumerGroupsExecuteFunc func(r ApiGetConsumerGroupsRequest) (ConsumerGroupList, *_nethttp.Response, error)

	// ResetConsumerGroupOffsetFunc mocks the ResetConsumerGroupOffset method.
	ResetConsumerGroupOffsetFunc func(ctx _context.Context, consumerGroupId string) ApiResetConsumerGroupOffsetRequest

	// ResetConsumerGroupOffsetExecuteFunc mocks the ResetConsumerGroupOffsetExecute method.
	ResetConsumerGroupOffsetExecuteFunc func(r ApiResetConsumerGroupOffsetRequest) (ConsumerGroupResetOffsetResult, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteConsumerGroupById holds details about calls to the DeleteConsumerGroupById method.
		DeleteConsumerGroupById []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ConsumerGroupId is the consumerGroupId argument value.
			ConsumerGroupId string
		}
		// DeleteConsumerGroupByIdExecute holds details about calls to the DeleteConsumerGroupByIdExecute method.
		DeleteConsumerGroupByIdExecute []struct {
			// R is the r argument value.
			R ApiDeleteConsumerGroupByIdRequest
		}
		// GetConsumerGroupById holds details about calls to the GetConsumerGroupById method.
		GetConsumerGroupById []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ConsumerGroupId is the consumerGroupId argument value.
			ConsumerGroupId string
		}
		// GetConsumerGroupByIdExecute holds details about calls to the GetConsumerGroupByIdExecute method.
		GetConsumerGroupByIdExecute []struct {
			// R is the r argument value.
			R ApiGetConsumerGroupByIdRequest
		}
		// GetConsumerGroups holds details about calls to the GetConsumerGroups method.
		GetConsumerGroups []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetConsumerGroupsExecute holds details about calls to the GetConsumerGroupsExecute method.
		GetConsumerGroupsExecute []struct {
			// R is the r argument value.
			R ApiGetConsumerGroupsRequest
		}
		// ResetConsumerGroupOffset holds details about calls to the ResetConsumerGroupOffset method.
		ResetConsumerGroupOffset []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// ConsumerGroupId is the consumerGroupId argument value.
			ConsumerGroupId string
		}
		// ResetConsumerGroupOffsetExecute holds details about calls to the ResetConsumerGroupOffsetExecute method.
		ResetConsumerGroupOffsetExecute []struct {
			// R is the r argument value.
			R ApiResetConsumerGroupOffsetRequest
		}
	}
	lockDeleteConsumerGroupById         sync.RWMutex
	lockDeleteConsumerGroupByIdExecute  sync.RWMutex
	lockGetConsumerGroupById            sync.RWMutex
	lockGetConsumerGroupByIdExecute     sync.RWMutex
	lockGetConsumerGroups               sync.RWMutex
	lockGetConsumerGroupsExecute        sync.RWMutex
	lockResetConsumerGroupOffset        sync.RWMutex
	lockResetConsumerGroupOffsetExecute sync.RWMutex
}

// DeleteConsumerGroupById calls DeleteConsumerGroupByIdFunc.
func (mock *GroupsApiMock) DeleteConsumerGroupById(ctx _context.Context, consumerGroupId string) ApiDeleteConsumerGroupByIdRequest {
	if mock.DeleteConsumerGroupByIdFunc == nil {
		panic("GroupsApiMock.DeleteConsumerGroupByIdFunc: method is nil but GroupsApi.DeleteConsumerGroupById was just called")
	}
	callInfo := struct {
		Ctx             _context.Context
		ConsumerGroupId string
	}{
		Ctx:             ctx,
		ConsumerGroupId: consumerGroupId,
	}
	mock.lockDeleteConsumerGroupById.Lock()
	mock.calls.DeleteConsumerGroupById = append(mock.calls.DeleteConsumerGroupById, callInfo)
	mock.lockDeleteConsumerGroupById.Unlock()
	return mock.DeleteConsumerGroupByIdFunc(ctx, consumerGroupId)
}

// DeleteConsumerGroupByIdCalls gets all the calls that were made to DeleteConsumerGroupById.
// Check the length with:
//     len(mockedGroupsApi.DeleteConsumerGroupByIdCalls())
func (mock *GroupsApiMock) DeleteConsumerGroupByIdCalls() []struct {
	Ctx             _context.Context
	ConsumerGroupId string
} {
	var calls []struct {
		Ctx             _context.Context
		ConsumerGroupId string
	}
	mock.lockDeleteConsumerGroupById.RLock()
	calls = mock.calls.DeleteConsumerGroupById
	mock.lockDeleteConsumerGroupById.RUnlock()
	return calls
}

// DeleteConsumerGroupByIdExecute calls DeleteConsumerGroupByIdExecuteFunc.
func (mock *GroupsApiMock) DeleteConsumerGroupByIdExecute(r ApiDeleteConsumerGroupByIdRequest) (*_nethttp.Response, error) {
	if mock.DeleteConsumerGroupByIdExecuteFunc == nil {
		panic("GroupsApiMock.DeleteConsumerGroupByIdExecuteFunc: method is nil but GroupsApi.DeleteConsumerGroupByIdExecute was just called")
	}
	callInfo := struct {
		R ApiDeleteConsumerGroupByIdRequest
	}{
		R: r,
	}
	mock.lockDeleteConsumerGroupByIdExecute.Lock()
	mock.calls.DeleteConsumerGroupByIdExecute = append(mock.calls.DeleteConsumerGroupByIdExecute, callInfo)
	mock.lockDeleteConsumerGroupByIdExecute.Unlock()
	return mock.DeleteConsumerGroupByIdExecuteFunc(r)
}

// DeleteConsumerGroupByIdExecuteCalls gets all the calls that were made to DeleteConsumerGroupByIdExecute.
// Check the length with:
//     len(mockedGroupsApi.DeleteConsumerGroupByIdExecuteCalls())
func (mock *GroupsApiMock) DeleteConsumerGroupByIdExecuteCalls() []struct {
	R ApiDeleteConsumerGroupByIdRequest
} {
	var calls []struct {
		R ApiDeleteConsumerGroupByIdRequest
	}
	mock.lockDeleteConsumerGroupByIdExecute.RLock()
	calls = mock.calls.DeleteConsumerGroupByIdExecute
	mock.lockDeleteConsumerGroupByIdExecute.RUnlock()
	return calls
}

// GetConsumerGroupById calls GetConsumerGroupByIdFunc.
func (mock *GroupsApiMock) GetConsumerGroupById(ctx _context.Context, consumerGroupId string) ApiGetConsumerGroupByIdRequest {
	if mock.GetConsumerGroupByIdFunc == nil {
		panic("GroupsApiMock.GetConsumerGroupByIdFunc: method is nil but GroupsApi.GetConsumerGroupById was just called")
	}
	callInfo := struct {
		Ctx             _context.Context
		ConsumerGroupId string
	}{
		Ctx:             ctx,
		ConsumerGroupId: consumerGroupId,
	}
	mock.lockGetConsumerGroupById.Lock()
	mock.calls.GetConsumerGroupById = append(mock.calls.GetConsumerGroupById, callInfo)
	mock.lockGetConsumerGroupById.Unlock()
	return mock.GetConsumerGroupByIdFunc(ctx, consumerGroupId)
}

// GetConsumerGroupByIdCalls gets all the calls that were made to GetConsumerGroupById.
// Check the length with:
//     len(mockedGroupsApi.GetConsumerGroupByIdCalls())
func (mock *GroupsApiMock) GetConsumerGroupByIdCalls() []struct {
	Ctx             _context.Context
	ConsumerGroupId string
} {
	var calls []struct {
		Ctx             _context.Context
		ConsumerGroupId string
	}
	mock.lockGetConsumerGroupById.RLock()
	calls = mock.calls.GetConsumerGroupById
	mock.lockGetConsumerGroupById.RUnlock()
	return calls
}

// GetConsumerGroupByIdExecute calls GetConsumerGroupByIdExecuteFunc.
func (mock *GroupsApiMock) GetConsumerGroupByIdExecute(r ApiGetConsumerGroupByIdRequest) (ConsumerGroup, *_nethttp.Response, error) {
	if mock.GetConsumerGroupByIdExecuteFunc == nil {
		panic("GroupsApiMock.GetConsumerGroupByIdExecuteFunc: method is nil but GroupsApi.GetConsumerGroupByIdExecute was just called")
	}
	callInfo := struct {
		R ApiGetConsumerGroupByIdRequest
	}{
		R: r,
	}
	mock.lockGetConsumerGroupByIdExecute.Lock()
	mock.calls.GetConsumerGroupByIdExecute = append(mock.calls.GetConsumerGroupByIdExecute, callInfo)
	mock.lockGetConsumerGroupByIdExecute.Unlock()
	return mock.GetConsumerGroupByIdExecuteFunc(r)
}

// GetConsumerGroupByIdExecuteCalls gets all the calls that were made to GetConsumerGroupByIdExecute.
// Check the length with:
//     len(mockedGroupsApi.GetConsumerGroupByIdExecuteCalls())
func (mock *GroupsApiMock) GetConsumerGroupByIdExecuteCalls() []struct {
	R ApiGetConsumerGroupByIdRequest
} {
	var calls []struct {
		R ApiGetConsumerGroupByIdRequest
	}
	mock.lockGetConsumerGroupByIdExecute.RLock()
	calls = mock.calls.GetConsumerGroupByIdExecute
	mock.lockGetConsumerGroupByIdExecute.RUnlock()
	return calls
}

// GetConsumerGroups calls GetConsumerGroupsFunc.
func (mock *GroupsApiMock) GetConsumerGroups(ctx _context.Context) ApiGetConsumerGroupsRequest {
	if mock.GetConsumerGroupsFunc == nil {
		panic("GroupsApiMock.GetConsumerGroupsFunc: method is nil but GroupsApi.GetConsumerGroups was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetConsumerGroups.Lock()
	mock.calls.GetConsumerGroups = append(mock.calls.GetConsumerGroups, callInfo)
	mock.lockGetConsumerGroups.Unlock()
	return mock.GetConsumerGroupsFunc(ctx)
}

// GetConsumerGroupsCalls gets all the calls that were made to GetConsumerGroups.
// Check the length with:
//     len(mockedGroupsApi.GetConsumerGroupsCalls())
func (mock *GroupsApiMock) GetConsumerGroupsCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetConsumerGroups.RLock()
	calls = mock.calls.GetConsumerGroups
	mock.lockGetConsumerGroups.RUnlock()
	return calls
}

// GetConsumerGroupsExecute calls GetConsumerGroupsExecuteFunc.
func (mock *GroupsApiMock) GetConsumerGroupsExecute(r ApiGetConsumerGroupsRequest) (ConsumerGroupList, *_nethttp.Response, error) {
	if mock.GetConsumerGroupsExecuteFunc == nil {
		panic("GroupsApiMock.GetConsumerGroupsExecuteFunc: method is nil but GroupsApi.GetConsumerGroupsExecute was just called")
	}
	callInfo := struct {
		R ApiGetConsumerGroupsRequest
	}{
		R: r,
	}
	mock.lockGetConsumerGroupsExecute.Lock()
	mock.calls.GetConsumerGroupsExecute = append(mock.calls.GetConsumerGroupsExecute, callInfo)
	mock.lockGetConsumerGroupsExecute.Unlock()
	return mock.GetConsumerGroupsExecuteFunc(r)
}

// GetConsumerGroupsExecuteCalls gets all the calls that were made to GetConsumerGroupsExecute.
// Check the length with:
//     len(mockedGroupsApi.GetConsumerGroupsExecuteCalls())
func (mock *GroupsApiMock) GetConsumerGroupsExecuteCalls() []struct {
	R ApiGetConsumerGroupsRequest
} {
	var calls []struct {
		R ApiGetConsumerGroupsRequest
	}
	mock.lockGetConsumerGroupsExecute.RLock()
	calls = mock.calls.GetConsumerGroupsExecute
	mock.lockGetConsumerGroupsExecute.RUnlock()
	return calls
}

// ResetConsumerGroupOffset calls ResetConsumerGroupOffsetFunc.
func (mock *GroupsApiMock) ResetConsumerGroupOffset(ctx _context.Context, consumerGroupId string) ApiResetConsumerGroupOffsetRequest {
	if mock.ResetConsumerGroupOffsetFunc == nil {
		panic("GroupsApiMock.ResetConsumerGroupOffsetFunc: method is nil but GroupsApi.ResetConsumerGroupOffset was just called")
	}
	callInfo := struct {
		Ctx             _context.Context
		ConsumerGroupId string
	}{
		Ctx:             ctx,
		ConsumerGroupId: consumerGroupId,
	}
	mock.lockResetConsumerGroupOffset.Lock()
	mock.calls.ResetConsumerGroupOffset = append(mock.calls.ResetConsumerGroupOffset, callInfo)
	mock.lockResetConsumerGroupOffset.Unlock()
	return mock.ResetConsumerGroupOffsetFunc(ctx, consumerGroupId)
}

// ResetConsumerGroupOffsetCalls gets all the calls that were made to ResetConsumerGroupOffset.
// Check the length with:
//     len(mockedGroupsApi.ResetConsumerGroupOffsetCalls())
func (mock *GroupsApiMock) ResetConsumerGroupOffsetCalls() []struct {
	Ctx             _context.Context
	ConsumerGroupId string
} {
	var calls []struct {
		Ctx             _context.Context
		ConsumerGroupId string
	}
	mock.lockResetConsumerGroupOffset.RLock()
	calls = mock.calls.ResetConsumerGroupOffset
	mock.lockResetConsumerGroupOffset.RUnlock()
	return calls
}

// ResetConsumerGroupOffsetExecute calls ResetConsumerGroupOffsetExecuteFunc.
func (mock *GroupsApiMock) ResetConsumerGroupOffsetExecute(r ApiResetConsumerGroupOffsetRequest) (ConsumerGroupResetOffsetResult, *_nethttp.Response, error) {
	if mock.ResetConsumerGroupOffsetExecuteFunc == nil {
		panic("GroupsApiMock.ResetConsumerGroupOffsetExecuteFunc: method is nil but GroupsApi.ResetConsumerGroupOffsetExecute was just called")
	}
	callInfo := struct {
		R ApiResetConsumerGroupOffsetRequest
	}{
		R: r,
	}
	mock.lockResetConsumerGroupOffsetExecute.Lock()
	mock.calls.ResetConsumerGroupOffsetExecute = append(mock.calls.ResetConsumerGroupOffsetExecute, callInfo)
	mock.lockResetConsumerGroupOffsetExecute.Unlock()
	return mock.ResetConsumerGroupOffsetExecuteFunc(r)
}

// ResetConsumerGroupOffsetExecuteCalls gets all the calls that were made to ResetConsumerGroupOffsetExecute.
// Check the length with:
//     len(mockedGroupsApi.ResetConsumerGroupOffsetExecuteCalls())
func (mock *GroupsApiMock) ResetConsumerGroupOffsetExecuteCalls() []struct {
	R ApiResetConsumerGroupOffsetRequest
} {
	var calls []struct {
		R ApiResetConsumerGroupOffsetRequest
	}
	mock.lockResetConsumerGroupOffsetExecute.RLock()
	calls = mock.calls.ResetConsumerGroupOffsetExecute
	mock.lockResetConsumerGroupOffsetExecute.RUnlock()
	return calls
}
