// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package kafkainstanceclient

import (
	_context "context"
	_nethttp "net/http"
	"sync"
)

// Ensure, that TopicsApiMock does implement TopicsApi.
// If this is not the case, regenerate this file with moq.
var _ TopicsApi = &TopicsApiMock{}

// TopicsApiMock is a mock implementation of TopicsApi.
//
// 	func TestSomethingThatUsesTopicsApi(t *testing.T) {
//
// 		// make and configure a mocked TopicsApi
// 		mockedTopicsApi := &TopicsApiMock{
// 			CreateTopicFunc: func(ctx _context.Context) ApiCreateTopicRequest {
// 				panic("mock out the CreateTopic method")
// 			},
// 			CreateTopicExecuteFunc: func(r ApiCreateTopicRequest) (Topic, *_nethttp.Response, error) {
// 				panic("mock out the CreateTopicExecute method")
// 			},
// 			DeleteTopicFunc: func(ctx _context.Context, topicName string) ApiDeleteTopicRequest {
// 				panic("mock out the DeleteTopic method")
// 			},
// 			DeleteTopicExecuteFunc: func(r ApiDeleteTopicRequest) (*_nethttp.Response, error) {
// 				panic("mock out the DeleteTopicExecute method")
// 			},
// 			GetTopicFunc: func(ctx _context.Context, topicName string) ApiGetTopicRequest {
// 				panic("mock out the GetTopic method")
// 			},
// 			GetTopicExecuteFunc: func(r ApiGetTopicRequest) (Topic, *_nethttp.Response, error) {
// 				panic("mock out the GetTopicExecute method")
// 			},
// 			GetTopicsFunc: func(ctx _context.Context) ApiGetTopicsRequest {
// 				panic("mock out the GetTopics method")
// 			},
// 			GetTopicsExecuteFunc: func(r ApiGetTopicsRequest) (TopicsList, *_nethttp.Response, error) {
// 				panic("mock out the GetTopicsExecute method")
// 			},
// 			UpdateTopicFunc: func(ctx _context.Context, topicName string) ApiUpdateTopicRequest {
// 				panic("mock out the UpdateTopic method")
// 			},
// 			UpdateTopicExecuteFunc: func(r ApiUpdateTopicRequest) (Topic, *_nethttp.Response, error) {
// 				panic("mock out the UpdateTopicExecute method")
// 			},
// 		}
//
// 		// use mockedTopicsApi in code that requires TopicsApi
// 		// and then make assertions.
//
// 	}
type TopicsApiMock struct {
	// CreateTopicFunc mocks the CreateTopic method.
	CreateTopicFunc func(ctx _context.Context) ApiCreateTopicRequest

	// CreateTopicExecuteFunc mocks the CreateTopicExecute method.
	CreateTopicExecuteFunc func(r ApiCreateTopicRequest) (Topic, *_nethttp.Response, error)

	// DeleteTopicFunc mocks the DeleteTopic method.
	DeleteTopicFunc func(ctx _context.Context, topicName string) ApiDeleteTopicRequest

	// DeleteTopicExecuteFunc mocks the DeleteTopicExecute method.
	DeleteTopicExecuteFunc func(r ApiDeleteTopicRequest) (*_nethttp.Response, error)

	// GetTopicFunc mocks the GetTopic method.
	GetTopicFunc func(ctx _context.Context, topicName string) ApiGetTopicRequest

	// GetTopicExecuteFunc mocks the GetTopicExecute method.
	GetTopicExecuteFunc func(r ApiGetTopicRequest) (Topic, *_nethttp.Response, error)

	// GetTopicsFunc mocks the GetTopics method.
	GetTopicsFunc func(ctx _context.Context) ApiGetTopicsRequest

	// GetTopicsExecuteFunc mocks the GetTopicsExecute method.
	GetTopicsExecuteFunc func(r ApiGetTopicsRequest) (TopicsList, *_nethttp.Response, error)

	// UpdateTopicFunc mocks the UpdateTopic method.
	UpdateTopicFunc func(ctx _context.Context, topicName string) ApiUpdateTopicRequest

	// UpdateTopicExecuteFunc mocks the UpdateTopicExecute method.
	UpdateTopicExecuteFunc func(r ApiUpdateTopicRequest) (Topic, *_nethttp.Response, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateTopic holds details about calls to the CreateTopic method.
		CreateTopic []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// CreateTopicExecute holds details about calls to the CreateTopicExecute method.
		CreateTopicExecute []struct {
			// R is the r argument value.
			R ApiCreateTopicRequest
		}
		// DeleteTopic holds details about calls to the DeleteTopic method.
		DeleteTopic []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// TopicName is the topicName argument value.
			TopicName string
		}
		// DeleteTopicExecute holds details about calls to the DeleteTopicExecute method.
		DeleteTopicExecute []struct {
			// R is the r argument value.
			R ApiDeleteTopicRequest
		}
		// GetTopic holds details about calls to the GetTopic method.
		GetTopic []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// TopicName is the topicName argument value.
			TopicName string
		}
		// GetTopicExecute holds details about calls to the GetTopicExecute method.
		GetTopicExecute []struct {
			// R is the r argument value.
			R ApiGetTopicRequest
		}
		// GetTopics holds details about calls to the GetTopics method.
		GetTopics []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
		}
		// GetTopicsExecute holds details about calls to the GetTopicsExecute method.
		GetTopicsExecute []struct {
			// R is the r argument value.
			R ApiGetTopicsRequest
		}
		// UpdateTopic holds details about calls to the UpdateTopic method.
		UpdateTopic []struct {
			// Ctx is the ctx argument value.
			Ctx _context.Context
			// TopicName is the topicName argument value.
			TopicName string
		}
		// UpdateTopicExecute holds details about calls to the UpdateTopicExecute method.
		UpdateTopicExecute []struct {
			// R is the r argument value.
			R ApiUpdateTopicRequest
		}
	}
	lockCreateTopic        sync.RWMutex
	lockCreateTopicExecute sync.RWMutex
	lockDeleteTopic        sync.RWMutex
	lockDeleteTopicExecute sync.RWMutex
	lockGetTopic           sync.RWMutex
	lockGetTopicExecute    sync.RWMutex
	lockGetTopics          sync.RWMutex
	lockGetTopicsExecute   sync.RWMutex
	lockUpdateTopic        sync.RWMutex
	lockUpdateTopicExecute sync.RWMutex
}

// CreateTopic calls CreateTopicFunc.
func (mock *TopicsApiMock) CreateTopic(ctx _context.Context) ApiCreateTopicRequest {
	if mock.CreateTopicFunc == nil {
		panic("TopicsApiMock.CreateTopicFunc: method is nil but TopicsApi.CreateTopic was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockCreateTopic.Lock()
	mock.calls.CreateTopic = append(mock.calls.CreateTopic, callInfo)
	mock.lockCreateTopic.Unlock()
	return mock.CreateTopicFunc(ctx)
}

// CreateTopicCalls gets all the calls that were made to CreateTopic.
// Check the length with:
//     len(mockedTopicsApi.CreateTopicCalls())
func (mock *TopicsApiMock) CreateTopicCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockCreateTopic.RLock()
	calls = mock.calls.CreateTopic
	mock.lockCreateTopic.RUnlock()
	return calls
}

// CreateTopicExecute calls CreateTopicExecuteFunc.
func (mock *TopicsApiMock) CreateTopicExecute(r ApiCreateTopicRequest) (Topic, *_nethttp.Response, error) {
	if mock.CreateTopicExecuteFunc == nil {
		panic("TopicsApiMock.CreateTopicExecuteFunc: method is nil but TopicsApi.CreateTopicExecute was just called")
	}
	callInfo := struct {
		R ApiCreateTopicRequest
	}{
		R: r,
	}
	mock.lockCreateTopicExecute.Lock()
	mock.calls.CreateTopicExecute = append(mock.calls.CreateTopicExecute, callInfo)
	mock.lockCreateTopicExecute.Unlock()
	return mock.CreateTopicExecuteFunc(r)
}

// CreateTopicExecuteCalls gets all the calls that were made to CreateTopicExecute.
// Check the length with:
//     len(mockedTopicsApi.CreateTopicExecuteCalls())
func (mock *TopicsApiMock) CreateTopicExecuteCalls() []struct {
	R ApiCreateTopicRequest
} {
	var calls []struct {
		R ApiCreateTopicRequest
	}
	mock.lockCreateTopicExecute.RLock()
	calls = mock.calls.CreateTopicExecute
	mock.lockCreateTopicExecute.RUnlock()
	return calls
}

// DeleteTopic calls DeleteTopicFunc.
func (mock *TopicsApiMock) DeleteTopic(ctx _context.Context, topicName string) ApiDeleteTopicRequest {
	if mock.DeleteTopicFunc == nil {
		panic("TopicsApiMock.DeleteTopicFunc: method is nil but TopicsApi.DeleteTopic was just called")
	}
	callInfo := struct {
		Ctx       _context.Context
		TopicName string
	}{
		Ctx:       ctx,
		TopicName: topicName,
	}
	mock.lockDeleteTopic.Lock()
	mock.calls.DeleteTopic = append(mock.calls.DeleteTopic, callInfo)
	mock.lockDeleteTopic.Unlock()
	return mock.DeleteTopicFunc(ctx, topicName)
}

// DeleteTopicCalls gets all the calls that were made to DeleteTopic.
// Check the length with:
//     len(mockedTopicsApi.DeleteTopicCalls())
func (mock *TopicsApiMock) DeleteTopicCalls() []struct {
	Ctx       _context.Context
	TopicName string
} {
	var calls []struct {
		Ctx       _context.Context
		TopicName string
	}
	mock.lockDeleteTopic.RLock()
	calls = mock.calls.DeleteTopic
	mock.lockDeleteTopic.RUnlock()
	return calls
}

// DeleteTopicExecute calls DeleteTopicExecuteFunc.
func (mock *TopicsApiMock) DeleteTopicExecute(r ApiDeleteTopicRequest) (*_nethttp.Response, error) {
	if mock.DeleteTopicExecuteFunc == nil {
		panic("TopicsApiMock.DeleteTopicExecuteFunc: method is nil but TopicsApi.DeleteTopicExecute was just called")
	}
	callInfo := struct {
		R ApiDeleteTopicRequest
	}{
		R: r,
	}
	mock.lockDeleteTopicExecute.Lock()
	mock.calls.DeleteTopicExecute = append(mock.calls.DeleteTopicExecute, callInfo)
	mock.lockDeleteTopicExecute.Unlock()
	return mock.DeleteTopicExecuteFunc(r)
}

// DeleteTopicExecuteCalls gets all the calls that were made to DeleteTopicExecute.
// Check the length with:
//     len(mockedTopicsApi.DeleteTopicExecuteCalls())
func (mock *TopicsApiMock) DeleteTopicExecuteCalls() []struct {
	R ApiDeleteTopicRequest
} {
	var calls []struct {
		R ApiDeleteTopicRequest
	}
	mock.lockDeleteTopicExecute.RLock()
	calls = mock.calls.DeleteTopicExecute
	mock.lockDeleteTopicExecute.RUnlock()
	return calls
}

// GetTopic calls GetTopicFunc.
func (mock *TopicsApiMock) GetTopic(ctx _context.Context, topicName string) ApiGetTopicRequest {
	if mock.GetTopicFunc == nil {
		panic("TopicsApiMock.GetTopicFunc: method is nil but TopicsApi.GetTopic was just called")
	}
	callInfo := struct {
		Ctx       _context.Context
		TopicName string
	}{
		Ctx:       ctx,
		TopicName: topicName,
	}
	mock.lockGetTopic.Lock()
	mock.calls.GetTopic = append(mock.calls.GetTopic, callInfo)
	mock.lockGetTopic.Unlock()
	return mock.GetTopicFunc(ctx, topicName)
}

// GetTopicCalls gets all the calls that were made to GetTopic.
// Check the length with:
//     len(mockedTopicsApi.GetTopicCalls())
func (mock *TopicsApiMock) GetTopicCalls() []struct {
	Ctx       _context.Context
	TopicName string
} {
	var calls []struct {
		Ctx       _context.Context
		TopicName string
	}
	mock.lockGetTopic.RLock()
	calls = mock.calls.GetTopic
	mock.lockGetTopic.RUnlock()
	return calls
}

// GetTopicExecute calls GetTopicExecuteFunc.
func (mock *TopicsApiMock) GetTopicExecute(r ApiGetTopicRequest) (Topic, *_nethttp.Response, error) {
	if mock.GetTopicExecuteFunc == nil {
		panic("TopicsApiMock.GetTopicExecuteFunc: method is nil but TopicsApi.GetTopicExecute was just called")
	}
	callInfo := struct {
		R ApiGetTopicRequest
	}{
		R: r,
	}
	mock.lockGetTopicExecute.Lock()
	mock.calls.GetTopicExecute = append(mock.calls.GetTopicExecute, callInfo)
	mock.lockGetTopicExecute.Unlock()
	return mock.GetTopicExecuteFunc(r)
}

// GetTopicExecuteCalls gets all the calls that were made to GetTopicExecute.
// Check the length with:
//     len(mockedTopicsApi.GetTopicExecuteCalls())
func (mock *TopicsApiMock) GetTopicExecuteCalls() []struct {
	R ApiGetTopicRequest
} {
	var calls []struct {
		R ApiGetTopicRequest
	}
	mock.lockGetTopicExecute.RLock()
	calls = mock.calls.GetTopicExecute
	mock.lockGetTopicExecute.RUnlock()
	return calls
}

// GetTopics calls GetTopicsFunc.
func (mock *TopicsApiMock) GetTopics(ctx _context.Context) ApiGetTopicsRequest {
	if mock.GetTopicsFunc == nil {
		panic("TopicsApiMock.GetTopicsFunc: method is nil but TopicsApi.GetTopics was just called")
	}
	callInfo := struct {
		Ctx _context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetTopics.Lock()
	mock.calls.GetTopics = append(mock.calls.GetTopics, callInfo)
	mock.lockGetTopics.Unlock()
	return mock.GetTopicsFunc(ctx)
}

// GetTopicsCalls gets all the calls that were made to GetTopics.
// Check the length with:
//     len(mockedTopicsApi.GetTopicsCalls())
func (mock *TopicsApiMock) GetTopicsCalls() []struct {
	Ctx _context.Context
} {
	var calls []struct {
		Ctx _context.Context
	}
	mock.lockGetTopics.RLock()
	calls = mock.calls.GetTopics
	mock.lockGetTopics.RUnlock()
	return calls
}

// GetTopicsExecute calls GetTopicsExecuteFunc.
func (mock *TopicsApiMock) GetTopicsExecute(r ApiGetTopicsRequest) (TopicsList, *_nethttp.Response, error) {
	if mock.GetTopicsExecuteFunc == nil {
		panic("TopicsApiMock.GetTopicsExecuteFunc: method is nil but TopicsApi.GetTopicsExecute was just called")
	}
	callInfo := struct {
		R ApiGetTopicsRequest
	}{
		R: r,
	}
	mock.lockGetTopicsExecute.Lock()
	mock.calls.GetTopicsExecute = append(mock.calls.GetTopicsExecute, callInfo)
	mock.lockGetTopicsExecute.Unlock()
	return mock.GetTopicsExecuteFunc(r)
}

// GetTopicsExecuteCalls gets all the calls that were made to GetTopicsExecute.
// Check the length with:
//     len(mockedTopicsApi.GetTopicsExecuteCalls())
func (mock *TopicsApiMock) GetTopicsExecuteCalls() []struct {
	R ApiGetTopicsRequest
} {
	var calls []struct {
		R ApiGetTopicsRequest
	}
	mock.lockGetTopicsExecute.RLock()
	calls = mock.calls.GetTopicsExecute
	mock.lockGetTopicsExecute.RUnlock()
	return calls
}

// UpdateTopic calls UpdateTopicFunc.
func (mock *TopicsApiMock) UpdateTopic(ctx _context.Context, topicName string) ApiUpdateTopicRequest {
	if mock.UpdateTopicFunc == nil {
		panic("TopicsApiMock.UpdateTopicFunc: method is nil but TopicsApi.UpdateTopic was just called")
	}
	callInfo := struct {
		Ctx       _context.Context
		TopicName string
	}{
		Ctx:       ctx,
		TopicName: topicName,
	}
	mock.lockUpdateTopic.Lock()
	mock.calls.UpdateTopic = append(mock.calls.UpdateTopic, callInfo)
	mock.lockUpdateTopic.Unlock()
	return mock.UpdateTopicFunc(ctx, topicName)
}

// UpdateTopicCalls gets all the calls that were made to UpdateTopic.
// Check the length with:
//     len(mockedTopicsApi.UpdateTopicCalls())
func (mock *TopicsApiMock) UpdateTopicCalls() []struct {
	Ctx       _context.Context
	TopicName string
} {
	var calls []struct {
		Ctx       _context.Context
		TopicName string
	}
	mock.lockUpdateTopic.RLock()
	calls = mock.calls.UpdateTopic
	mock.lockUpdateTopic.RUnlock()
	return calls
}

// UpdateTopicExecute calls UpdateTopicExecuteFunc.
func (mock *TopicsApiMock) UpdateTopicExecute(r ApiUpdateTopicRequest) (Topic, *_nethttp.Response, error) {
	if mock.UpdateTopicExecuteFunc == nil {
		panic("TopicsApiMock.UpdateTopicExecuteFunc: method is nil but TopicsApi.UpdateTopicExecute was just called")
	}
	callInfo := struct {
		R ApiUpdateTopicRequest
	}{
		R: r,
	}
	mock.lockUpdateTopicExecute.Lock()
	mock.calls.UpdateTopicExecute = append(mock.calls.UpdateTopicExecute, callInfo)
	mock.lockUpdateTopicExecute.Unlock()
	return mock.UpdateTopicExecuteFunc(r)
}

// UpdateTopicExecuteCalls gets all the calls that were made to UpdateTopicExecute.
// Check the length with:
//     len(mockedTopicsApi.UpdateTopicExecuteCalls())
func (mock *TopicsApiMock) UpdateTopicExecuteCalls() []struct {
	R ApiUpdateTopicRequest
} {
	var calls []struct {
		R ApiUpdateTopicRequest
	}
	mock.lockUpdateTopicExecute.RLock()
	calls = mock.calls.UpdateTopicExecute
	mock.lockUpdateTopicExecute.RUnlock()
	return calls
}
