// Code generated by counterfeiter. DO NOT EDIT.
package routingfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/routing"
	"github.com/livekit/protocol/livekit"
	"google.golang.org/protobuf/proto"
)

type FakeMessageSource struct {
	CloseStub        func()
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
	}
	ConnectionIDStub        func() livekit.ConnectionID
	connectionIDMutex       sync.RWMutex
	connectionIDArgsForCall []struct {
	}
	connectionIDReturns struct {
		result1 livekit.ConnectionID
	}
	connectionIDReturnsOnCall map[int]struct {
		result1 livekit.ConnectionID
	}
	IsClosedStub        func() bool
	isClosedMutex       sync.RWMutex
	isClosedArgsForCall []struct {
	}
	isClosedReturns struct {
		result1 bool
	}
	isClosedReturnsOnCall map[int]struct {
		result1 bool
	}
	ReadChanStub        func() <-chan proto.Message
	readChanMutex       sync.RWMutex
	readChanArgsForCall []struct {
	}
	readChanReturns struct {
		result1 <-chan proto.Message
	}
	readChanReturnsOnCall map[int]struct {
		result1 <-chan proto.Message
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMessageSource) Close() {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
	}{})
	stub := fake.CloseStub
	fake.recordInvocation("Close", []interface{}{})
	fake.closeMutex.Unlock()
	if stub != nil {
		fake.CloseStub()
	}
}

func (fake *FakeMessageSource) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeMessageSource) CloseCalls(stub func()) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeMessageSource) ConnectionID() livekit.ConnectionID {
	fake.connectionIDMutex.Lock()
	ret, specificReturn := fake.connectionIDReturnsOnCall[len(fake.connectionIDArgsForCall)]
	fake.connectionIDArgsForCall = append(fake.connectionIDArgsForCall, struct {
	}{})
	stub := fake.ConnectionIDStub
	fakeReturns := fake.connectionIDReturns
	fake.recordInvocation("ConnectionID", []interface{}{})
	fake.connectionIDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMessageSource) ConnectionIDCallCount() int {
	fake.connectionIDMutex.RLock()
	defer fake.connectionIDMutex.RUnlock()
	return len(fake.connectionIDArgsForCall)
}

func (fake *FakeMessageSource) ConnectionIDCalls(stub func() livekit.ConnectionID) {
	fake.connectionIDMutex.Lock()
	defer fake.connectionIDMutex.Unlock()
	fake.ConnectionIDStub = stub
}

func (fake *FakeMessageSource) ConnectionIDReturns(result1 livekit.ConnectionID) {
	fake.connectionIDMutex.Lock()
	defer fake.connectionIDMutex.Unlock()
	fake.ConnectionIDStub = nil
	fake.connectionIDReturns = struct {
		result1 livekit.ConnectionID
	}{result1}
}

func (fake *FakeMessageSource) ConnectionIDReturnsOnCall(i int, result1 livekit.ConnectionID) {
	fake.connectionIDMutex.Lock()
	defer fake.connectionIDMutex.Unlock()
	fake.ConnectionIDStub = nil
	if fake.connectionIDReturnsOnCall == nil {
		fake.connectionIDReturnsOnCall = make(map[int]struct {
			result1 livekit.ConnectionID
		})
	}
	fake.connectionIDReturnsOnCall[i] = struct {
		result1 livekit.ConnectionID
	}{result1}
}

func (fake *FakeMessageSource) IsClosed() bool {
	fake.isClosedMutex.Lock()
	ret, specificReturn := fake.isClosedReturnsOnCall[len(fake.isClosedArgsForCall)]
	fake.isClosedArgsForCall = append(fake.isClosedArgsForCall, struct {
	}{})
	stub := fake.IsClosedStub
	fakeReturns := fake.isClosedReturns
	fake.recordInvocation("IsClosed", []interface{}{})
	fake.isClosedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMessageSource) IsClosedCallCount() int {
	fake.isClosedMutex.RLock()
	defer fake.isClosedMutex.RUnlock()
	return len(fake.isClosedArgsForCall)
}

func (fake *FakeMessageSource) IsClosedCalls(stub func() bool) {
	fake.isClosedMutex.Lock()
	defer fake.isClosedMutex.Unlock()
	fake.IsClosedStub = stub
}

func (fake *FakeMessageSource) IsClosedReturns(result1 bool) {
	fake.isClosedMutex.Lock()
	defer fake.isClosedMutex.Unlock()
	fake.IsClosedStub = nil
	fake.isClosedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMessageSource) IsClosedReturnsOnCall(i int, result1 bool) {
	fake.isClosedMutex.Lock()
	defer fake.isClosedMutex.Unlock()
	fake.IsClosedStub = nil
	if fake.isClosedReturnsOnCall == nil {
		fake.isClosedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isClosedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMessageSource) ReadChan() <-chan proto.Message {
	fake.readChanMutex.Lock()
	ret, specificReturn := fake.readChanReturnsOnCall[len(fake.readChanArgsForCall)]
	fake.readChanArgsForCall = append(fake.readChanArgsForCall, struct {
	}{})
	stub := fake.ReadChanStub
	fakeReturns := fake.readChanReturns
	fake.recordInvocation("ReadChan", []interface{}{})
	fake.readChanMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMessageSource) ReadChanCallCount() int {
	fake.readChanMutex.RLock()
	defer fake.readChanMutex.RUnlock()
	return len(fake.readChanArgsForCall)
}

func (fake *FakeMessageSource) ReadChanCalls(stub func() <-chan proto.Message) {
	fake.readChanMutex.Lock()
	defer fake.readChanMutex.Unlock()
	fake.ReadChanStub = stub
}

func (fake *FakeMessageSource) ReadChanReturns(result1 <-chan proto.Message) {
	fake.readChanMutex.Lock()
	defer fake.readChanMutex.Unlock()
	fake.ReadChanStub = nil
	fake.readChanReturns = struct {
		result1 <-chan proto.Message
	}{result1}
}

func (fake *FakeMessageSource) ReadChanReturnsOnCall(i int, result1 <-chan proto.Message) {
	fake.readChanMutex.Lock()
	defer fake.readChanMutex.Unlock()
	fake.ReadChanStub = nil
	if fake.readChanReturnsOnCall == nil {
		fake.readChanReturnsOnCall = make(map[int]struct {
			result1 <-chan proto.Message
		})
	}
	fake.readChanReturnsOnCall[i] = struct {
		result1 <-chan proto.Message
	}{result1}
}

func (fake *FakeMessageSource) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.connectionIDMutex.RLock()
	defer fake.connectionIDMutex.RUnlock()
	fake.isClosedMutex.RLock()
	defer fake.isClosedMutex.RUnlock()
	fake.readChanMutex.RLock()
	defer fake.readChanMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMessageSource) recordInvocation(key string, args []interface{}) {
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

var _ routing.MessageSource = new(FakeMessageSource)
