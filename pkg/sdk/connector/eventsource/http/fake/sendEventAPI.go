// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"github.com/keptn/go-utils/pkg/api/models"
	"sync"
)

// SendEventAPIMock is a mock implementation of http.SendEventAPI.
//
// 	func TestSomethingThatUsesSendEventAPI(t *testing.T) {
//
// 		// make and configure a mocked http.SendEventAPI
// 		mockedSendEventAPI := &SendEventAPIMock{
// 			SendEventFunc: func(keptnContextExtendedCE models.KeptnContextExtendedCE) (*models.EventContext, *models.Error) {
// 				panic("mock out the SendEvent method")
// 			},
// 		}
//
// 		// use mockedSendEventAPI in code that requires http.SendEventAPI
// 		// and then make assertions.
//
// 	}
type SendEventAPIMock struct {
	// SendEventFunc mocks the SendEvent method.
	SendEventFunc func(keptnContextExtendedCE models.KeptnContextExtendedCE) (*models.EventContext, *models.Error)

	// calls tracks calls to the methods.
	calls struct {
		// SendEvent holds details about calls to the SendEvent method.
		SendEvent []struct {
			// KeptnContextExtendedCE is the keptnContextExtendedCE argument value.
			KeptnContextExtendedCE models.KeptnContextExtendedCE
		}
	}
	lockSendEvent sync.RWMutex
}

// SendEvent calls SendEventFunc.
func (mock *SendEventAPIMock) SendEvent(keptnContextExtendedCE models.KeptnContextExtendedCE) (*models.EventContext, *models.Error) {
	if mock.SendEventFunc == nil {
		panic("SendEventAPIMock.SendEventFunc: method is nil but SendEventAPI.SendEvent was just called")
	}
	callInfo := struct {
		KeptnContextExtendedCE models.KeptnContextExtendedCE
	}{
		KeptnContextExtendedCE: keptnContextExtendedCE,
	}
	mock.lockSendEvent.Lock()
	mock.calls.SendEvent = append(mock.calls.SendEvent, callInfo)
	mock.lockSendEvent.Unlock()
	return mock.SendEventFunc(keptnContextExtendedCE)
}

// SendEventCalls gets all the calls that were made to SendEvent.
// Check the length with:
//     len(mockedSendEventAPI.SendEventCalls())
func (mock *SendEventAPIMock) SendEventCalls() []struct {
	KeptnContextExtendedCE models.KeptnContextExtendedCE
} {
	var calls []struct {
		KeptnContextExtendedCE models.KeptnContextExtendedCE
	}
	mock.lockSendEvent.RLock()
	calls = mock.calls.SendEvent
	mock.lockSendEvent.RUnlock()
	return calls
}
