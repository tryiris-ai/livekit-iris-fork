// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"
	"time"

	"github.com/livekit/livekit-server/pkg/rtc/types"
)

type FakeWebsocketClient struct {
	CloseStub        func() error
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	closeReturns struct {
		result1 error
	}
	closeReturnsOnCall map[int]struct {
		result1 error
	}
	ReadMessageStub        func() (int, []byte, error)
	readMessageMutex       sync.RWMutex
	readMessageArgsForCall []struct {
	}
	readMessageReturns struct {
		result1 int
		result2 []byte
		result3 error
	}
	readMessageReturnsOnCall map[int]struct {
		result1 int
		result2 []byte
		result3 error
	}
	SetReadDeadlineStub        func(time.Time) error
	setReadDeadlineMutex       sync.RWMutex
	setReadDeadlineArgsForCall []struct {
		arg1 time.Time
	}
	setReadDeadlineReturns struct {
		result1 error
	}
	setReadDeadlineReturnsOnCall map[int]struct {
		result1 error
	}
	WriteControlStub        func(int, []byte, time.Time) error
	writeControlMutex       sync.RWMutex
	writeControlArgsForCall []struct {
		arg1 int
		arg2 []byte
		arg3 time.Time
	}
	writeControlReturns struct {
		result1 error
	}
	writeControlReturnsOnCall map[int]struct {
		result1 error
	}
	WriteMessageStub        func(int, []byte) error
	writeMessageMutex       sync.RWMutex
	writeMessageArgsForCall []struct {
		arg1 int
		arg2 []byte
	}
	writeMessageReturns struct {
		result1 error
	}
	writeMessageReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWebsocketClient) Close() error {
	fake.closeMutex.Lock()
	ret, specificReturn := fake.closeReturnsOnCall[len(fake.closeArgsForCall)]
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fakeReturns := fake.closeReturns
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWebsocketClient) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeWebsocketClient) CloseCalls(stub func() error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeWebsocketClient) CloseReturns(result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	fake.closeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) CloseReturnsOnCall(i int, result1 error) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = nil
	if fake.closeReturnsOnCall == nil {
		fake.closeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.closeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) ReadMessage() (int, []byte, error) {
	fake.readMessageMutex.Lock()
	ret, specificReturn := fake.readMessageReturnsOnCall[len(fake.readMessageArgsForCall)]
	fake.readMessageArgsForCall = append(fake.readMessageArgsForCall, struct {
	}{})
	stub := fake.ReadMessageStub
	fakeReturns := fake.readMessageReturns
	fake.recordInvocation("ReadMessage", []interface{}{})
	fake.readMessageMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeWebsocketClient) ReadMessageCallCount() int {
	fake.readMessageMutex.RLock()
	defer fake.readMessageMutex.RUnlock()
	return len(fake.readMessageArgsForCall)
}

func (fake *FakeWebsocketClient) ReadMessageCalls(stub func() (int, []byte, error)) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = stub
}

func (fake *FakeWebsocketClient) ReadMessageReturns(result1 int, result2 []byte, result3 error) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = nil
	fake.readMessageReturns = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWebsocketClient) ReadMessageReturnsOnCall(i int, result1 int, result2 []byte, result3 error) {
	fake.readMessageMutex.Lock()
	defer fake.readMessageMutex.Unlock()
	fake.ReadMessageStub = nil
	if fake.readMessageReturnsOnCall == nil {
		fake.readMessageReturnsOnCall = make(map[int]struct {
			result1 int
			result2 []byte
			result3 error
		})
	}
	fake.readMessageReturnsOnCall[i] = struct {
		result1 int
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWebsocketClient) SetReadDeadline(arg1 time.Time) error {
	fake.setReadDeadlineMutex.Lock()
	ret, specificReturn := fake.setReadDeadlineReturnsOnCall[len(fake.setReadDeadlineArgsForCall)]
	fake.setReadDeadlineArgsForCall = append(fake.setReadDeadlineArgsForCall, struct {
		arg1 time.Time
	}{arg1})
	stub := fake.SetReadDeadlineStub
	fakeReturns := fake.setReadDeadlineReturns
	fake.recordInvocation("SetReadDeadline", []interface{}{arg1})
	fake.setReadDeadlineMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWebsocketClient) SetReadDeadlineCallCount() int {
	fake.setReadDeadlineMutex.RLock()
	defer fake.setReadDeadlineMutex.RUnlock()
	return len(fake.setReadDeadlineArgsForCall)
}

func (fake *FakeWebsocketClient) SetReadDeadlineCalls(stub func(time.Time) error) {
	fake.setReadDeadlineMutex.Lock()
	defer fake.setReadDeadlineMutex.Unlock()
	fake.SetReadDeadlineStub = stub
}

func (fake *FakeWebsocketClient) SetReadDeadlineArgsForCall(i int) time.Time {
	fake.setReadDeadlineMutex.RLock()
	defer fake.setReadDeadlineMutex.RUnlock()
	argsForCall := fake.setReadDeadlineArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeWebsocketClient) SetReadDeadlineReturns(result1 error) {
	fake.setReadDeadlineMutex.Lock()
	defer fake.setReadDeadlineMutex.Unlock()
	fake.SetReadDeadlineStub = nil
	fake.setReadDeadlineReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) SetReadDeadlineReturnsOnCall(i int, result1 error) {
	fake.setReadDeadlineMutex.Lock()
	defer fake.setReadDeadlineMutex.Unlock()
	fake.SetReadDeadlineStub = nil
	if fake.setReadDeadlineReturnsOnCall == nil {
		fake.setReadDeadlineReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setReadDeadlineReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) WriteControl(arg1 int, arg2 []byte, arg3 time.Time) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeControlMutex.Lock()
	ret, specificReturn := fake.writeControlReturnsOnCall[len(fake.writeControlArgsForCall)]
	fake.writeControlArgsForCall = append(fake.writeControlArgsForCall, struct {
		arg1 int
		arg2 []byte
		arg3 time.Time
	}{arg1, arg2Copy, arg3})
	stub := fake.WriteControlStub
	fakeReturns := fake.writeControlReturns
	fake.recordInvocation("WriteControl", []interface{}{arg1, arg2Copy, arg3})
	fake.writeControlMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWebsocketClient) WriteControlCallCount() int {
	fake.writeControlMutex.RLock()
	defer fake.writeControlMutex.RUnlock()
	return len(fake.writeControlArgsForCall)
}

func (fake *FakeWebsocketClient) WriteControlCalls(stub func(int, []byte, time.Time) error) {
	fake.writeControlMutex.Lock()
	defer fake.writeControlMutex.Unlock()
	fake.WriteControlStub = stub
}

func (fake *FakeWebsocketClient) WriteControlArgsForCall(i int) (int, []byte, time.Time) {
	fake.writeControlMutex.RLock()
	defer fake.writeControlMutex.RUnlock()
	argsForCall := fake.writeControlArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeWebsocketClient) WriteControlReturns(result1 error) {
	fake.writeControlMutex.Lock()
	defer fake.writeControlMutex.Unlock()
	fake.WriteControlStub = nil
	fake.writeControlReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) WriteControlReturnsOnCall(i int, result1 error) {
	fake.writeControlMutex.Lock()
	defer fake.writeControlMutex.Unlock()
	fake.WriteControlStub = nil
	if fake.writeControlReturnsOnCall == nil {
		fake.writeControlReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeControlReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) WriteMessage(arg1 int, arg2 []byte) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeMessageMutex.Lock()
	ret, specificReturn := fake.writeMessageReturnsOnCall[len(fake.writeMessageArgsForCall)]
	fake.writeMessageArgsForCall = append(fake.writeMessageArgsForCall, struct {
		arg1 int
		arg2 []byte
	}{arg1, arg2Copy})
	stub := fake.WriteMessageStub
	fakeReturns := fake.writeMessageReturns
	fake.recordInvocation("WriteMessage", []interface{}{arg1, arg2Copy})
	fake.writeMessageMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeWebsocketClient) WriteMessageCallCount() int {
	fake.writeMessageMutex.RLock()
	defer fake.writeMessageMutex.RUnlock()
	return len(fake.writeMessageArgsForCall)
}

func (fake *FakeWebsocketClient) WriteMessageCalls(stub func(int, []byte) error) {
	fake.writeMessageMutex.Lock()
	defer fake.writeMessageMutex.Unlock()
	fake.WriteMessageStub = stub
}

func (fake *FakeWebsocketClient) WriteMessageArgsForCall(i int) (int, []byte) {
	fake.writeMessageMutex.RLock()
	defer fake.writeMessageMutex.RUnlock()
	argsForCall := fake.writeMessageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeWebsocketClient) WriteMessageReturns(result1 error) {
	fake.writeMessageMutex.Lock()
	defer fake.writeMessageMutex.Unlock()
	fake.WriteMessageStub = nil
	fake.writeMessageReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) WriteMessageReturnsOnCall(i int, result1 error) {
	fake.writeMessageMutex.Lock()
	defer fake.writeMessageMutex.Unlock()
	fake.WriteMessageStub = nil
	if fake.writeMessageReturnsOnCall == nil {
		fake.writeMessageReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeMessageReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeWebsocketClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.readMessageMutex.RLock()
	defer fake.readMessageMutex.RUnlock()
	fake.setReadDeadlineMutex.RLock()
	defer fake.setReadDeadlineMutex.RUnlock()
	fake.writeControlMutex.RLock()
	defer fake.writeControlMutex.RUnlock()
	fake.writeMessageMutex.RLock()
	defer fake.writeMessageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeWebsocketClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ types.WebsocketClient = new(FakeWebsocketClient)
