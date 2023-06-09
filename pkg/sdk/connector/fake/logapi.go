// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"context"
	"github.com/Mouhamadou305/go-utils2/pkg/api/models"
	"sync"
)

// LogAPIMock is a mock implementation of controlplane.logAPI.
//
// 	func TestSomethingThatUseslogAPI(t *testing.T) {
//
// 		// make and configure a mocked controlplane.logAPI
// 		mockedlogAPI := &LogAPIMock{
// 			DeleteLogsFunc: func(filter models.LogFilter) error {
// 				panic("mock out the DeleteLogs method")
// 			},
// 			FlushFunc: func() error {
// 				panic("mock out the Flush method")
// 			},
// 			GetLogsFunc: func(params models.GetLogsParams) (*models.GetLogsResponse, error) {
// 				panic("mock out the GetLogs method")
// 			},
// 			LogFunc: func(logs []models.LogEntry)  {
// 				panic("mock out the Log method")
// 			},
// 			StartFunc: func(ctx context.Context)  {
// 				panic("mock out the Start method")
// 			},
// 		}
//
// 		// use mockedlogAPI in code that requires controlplane.logAPI
// 		// and then make assertions.
//
// 	}
type LogAPIMock struct {
	// DeleteLogsFunc mocks the DeleteLogs method.
	DeleteLogsFunc func(filter models.LogFilter) error

	// FlushFunc mocks the Flush method.
	FlushFunc func() error

	// GetLogsFunc mocks the GetLogs method.
	GetLogsFunc func(params models.GetLogsParams) (*models.GetLogsResponse, error)

	// LogFunc mocks the Log method.
	LogFunc func(logs []models.LogEntry)

	// StartFunc mocks the Start method.
	StartFunc func(ctx context.Context)

	// calls tracks calls to the methods.
	calls struct {
		// DeleteLogs holds details about calls to the DeleteLogs method.
		DeleteLogs []struct {
			// Filter is the filter argument value.
			Filter models.LogFilter
		}
		// Flush holds details about calls to the Flush method.
		Flush []struct {
		}
		// GetLogs holds details about calls to the GetLogs method.
		GetLogs []struct {
			// Params is the params argument value.
			Params models.GetLogsParams
		}
		// Log holds details about calls to the Log method.
		Log []struct {
			// Logs is the logs argument value.
			Logs []models.LogEntry
		}
		// Start holds details about calls to the Start method.
		Start []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockDeleteLogs sync.RWMutex
	lockFlush      sync.RWMutex
	lockGetLogs    sync.RWMutex
	lockLog        sync.RWMutex
	lockStart      sync.RWMutex
}

// DeleteLogs calls DeleteLogsFunc.
func (mock *LogAPIMock) DeleteLogs(filter models.LogFilter) error {
	if mock.DeleteLogsFunc == nil {
		panic("LogAPIMock.DeleteLogsFunc: method is nil but logAPI.DeleteLogs was just called")
	}
	callInfo := struct {
		Filter models.LogFilter
	}{
		Filter: filter,
	}
	mock.lockDeleteLogs.Lock()
	mock.calls.DeleteLogs = append(mock.calls.DeleteLogs, callInfo)
	mock.lockDeleteLogs.Unlock()
	return mock.DeleteLogsFunc(filter)
}

// DeleteLogsCalls gets all the calls that were made to DeleteLogs.
// Check the length with:
//     len(mockedlogAPI.DeleteLogsCalls())
func (mock *LogAPIMock) DeleteLogsCalls() []struct {
	Filter models.LogFilter
} {
	var calls []struct {
		Filter models.LogFilter
	}
	mock.lockDeleteLogs.RLock()
	calls = mock.calls.DeleteLogs
	mock.lockDeleteLogs.RUnlock()
	return calls
}

// Flush calls FlushFunc.
func (mock *LogAPIMock) Flush() error {
	if mock.FlushFunc == nil {
		panic("LogAPIMock.FlushFunc: method is nil but logAPI.Flush was just called")
	}
	callInfo := struct {
	}{}
	mock.lockFlush.Lock()
	mock.calls.Flush = append(mock.calls.Flush, callInfo)
	mock.lockFlush.Unlock()
	return mock.FlushFunc()
}

// FlushCalls gets all the calls that were made to Flush.
// Check the length with:
//     len(mockedlogAPI.FlushCalls())
func (mock *LogAPIMock) FlushCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockFlush.RLock()
	calls = mock.calls.Flush
	mock.lockFlush.RUnlock()
	return calls
}

// GetLogs calls GetLogsFunc.
func (mock *LogAPIMock) GetLogs(params models.GetLogsParams) (*models.GetLogsResponse, error) {
	if mock.GetLogsFunc == nil {
		panic("LogAPIMock.GetLogsFunc: method is nil but logAPI.GetLogs was just called")
	}
	callInfo := struct {
		Params models.GetLogsParams
	}{
		Params: params,
	}
	mock.lockGetLogs.Lock()
	mock.calls.GetLogs = append(mock.calls.GetLogs, callInfo)
	mock.lockGetLogs.Unlock()
	return mock.GetLogsFunc(params)
}

// GetLogsCalls gets all the calls that were made to GetLogs.
// Check the length with:
//     len(mockedlogAPI.GetLogsCalls())
func (mock *LogAPIMock) GetLogsCalls() []struct {
	Params models.GetLogsParams
} {
	var calls []struct {
		Params models.GetLogsParams
	}
	mock.lockGetLogs.RLock()
	calls = mock.calls.GetLogs
	mock.lockGetLogs.RUnlock()
	return calls
}

// Log calls LogFunc.
func (mock *LogAPIMock) Log(logs []models.LogEntry) {
	if mock.LogFunc == nil {
		panic("LogAPIMock.LogFunc: method is nil but logAPI.Log was just called")
	}
	callInfo := struct {
		Logs []models.LogEntry
	}{
		Logs: logs,
	}
	mock.lockLog.Lock()
	mock.calls.Log = append(mock.calls.Log, callInfo)
	mock.lockLog.Unlock()
	mock.LogFunc(logs)
}

// LogCalls gets all the calls that were made to Log.
// Check the length with:
//     len(mockedlogAPI.LogCalls())
func (mock *LogAPIMock) LogCalls() []struct {
	Logs []models.LogEntry
} {
	var calls []struct {
		Logs []models.LogEntry
	}
	mock.lockLog.RLock()
	calls = mock.calls.Log
	mock.lockLog.RUnlock()
	return calls
}

// Start calls StartFunc.
func (mock *LogAPIMock) Start(ctx context.Context) {
	if mock.StartFunc == nil {
		panic("LogAPIMock.StartFunc: method is nil but logAPI.Start was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockStart.Lock()
	mock.calls.Start = append(mock.calls.Start, callInfo)
	mock.lockStart.Unlock()
	mock.StartFunc(ctx)
}

// StartCalls gets all the calls that were made to Start.
// Check the length with:
//     len(mockedlogAPI.StartCalls())
func (mock *LogAPIMock) StartCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockStart.RLock()
	calls = mock.calls.Start
	mock.lockStart.RUnlock()
	return calls
}
