// Code generated by counterfeiter. DO NOT EDIT.
package typesfakes

import (
	"sync"

	"github.com/livekit/livekit-server/pkg/rtc/types"
	"github.com/livekit/protocol/livekit"
)

type FakeMediaTrack struct {
	AddSubscriberStub        func(types.Participant) error
	addSubscriberMutex       sync.RWMutex
	addSubscriberArgsForCall []struct {
		arg1 types.Participant
	}
	addSubscriberReturns struct {
		result1 error
	}
	addSubscriberReturnsOnCall map[int]struct {
		result1 error
	}
	GetQualityForDimensionStub        func(uint32, uint32) livekit.VideoQuality
	getQualityForDimensionMutex       sync.RWMutex
	getQualityForDimensionArgsForCall []struct {
		arg1 uint32
		arg2 uint32
	}
	getQualityForDimensionReturns struct {
		result1 livekit.VideoQuality
	}
	getQualityForDimensionReturnsOnCall map[int]struct {
		result1 livekit.VideoQuality
	}
	IDStub        func() string
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 string
	}
	iDReturnsOnCall map[int]struct {
		result1 string
	}
	IsMutedStub        func() bool
	isMutedMutex       sync.RWMutex
	isMutedArgsForCall []struct {
	}
	isMutedReturns struct {
		result1 bool
	}
	isMutedReturnsOnCall map[int]struct {
		result1 bool
	}
	IsSimulcastStub        func() bool
	isSimulcastMutex       sync.RWMutex
	isSimulcastArgsForCall []struct {
	}
	isSimulcastReturns struct {
		result1 bool
	}
	isSimulcastReturnsOnCall map[int]struct {
		result1 bool
	}
	IsSubscriberStub        func(string) bool
	isSubscriberMutex       sync.RWMutex
	isSubscriberArgsForCall []struct {
		arg1 string
	}
	isSubscriberReturns struct {
		result1 bool
	}
	isSubscriberReturnsOnCall map[int]struct {
		result1 bool
	}
	KindStub        func() livekit.TrackType
	kindMutex       sync.RWMutex
	kindArgsForCall []struct {
	}
	kindReturns struct {
		result1 livekit.TrackType
	}
	kindReturnsOnCall map[int]struct {
		result1 livekit.TrackType
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	NotifySubscriberMaxQualityStub        func(string, livekit.VideoQuality)
	notifySubscriberMaxQualityMutex       sync.RWMutex
	notifySubscriberMaxQualityArgsForCall []struct {
		arg1 string
		arg2 livekit.VideoQuality
	}
	NotifySubscriberMuteStub        func(string)
	notifySubscriberMuteMutex       sync.RWMutex
	notifySubscriberMuteArgsForCall []struct {
		arg1 string
	}
	OnSubscribedMaxQualityChangeStub        func(func(trackSid string, subscribedQualities []*livekit.SubscribedQuality) error)
	onSubscribedMaxQualityChangeMutex       sync.RWMutex
	onSubscribedMaxQualityChangeArgsForCall []struct {
		arg1 func(trackSid string, subscribedQualities []*livekit.SubscribedQuality) error
	}
	RemoveAllSubscribersStub        func()
	removeAllSubscribersMutex       sync.RWMutex
	removeAllSubscribersArgsForCall []struct {
	}
	RemoveSubscriberStub        func(string)
	removeSubscriberMutex       sync.RWMutex
	removeSubscriberArgsForCall []struct {
		arg1 string
	}
	SetMutedStub        func(bool)
	setMutedMutex       sync.RWMutex
	setMutedArgsForCall []struct {
		arg1 bool
	}
	SourceStub        func() livekit.TrackSource
	sourceMutex       sync.RWMutex
	sourceArgsForCall []struct {
	}
	sourceReturns struct {
		result1 livekit.TrackSource
	}
	sourceReturnsOnCall map[int]struct {
		result1 livekit.TrackSource
	}
	UpdateVideoLayersStub        func([]*livekit.VideoLayer)
	updateVideoLayersMutex       sync.RWMutex
	updateVideoLayersArgsForCall []struct {
		arg1 []*livekit.VideoLayer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMediaTrack) AddSubscriber(arg1 types.Participant) error {
	fake.addSubscriberMutex.Lock()
	ret, specificReturn := fake.addSubscriberReturnsOnCall[len(fake.addSubscriberArgsForCall)]
	fake.addSubscriberArgsForCall = append(fake.addSubscriberArgsForCall, struct {
		arg1 types.Participant
	}{arg1})
	stub := fake.AddSubscriberStub
	fakeReturns := fake.addSubscriberReturns
	fake.recordInvocation("AddSubscriber", []interface{}{arg1})
	fake.addSubscriberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) AddSubscriberCallCount() int {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	return len(fake.addSubscriberArgsForCall)
}

func (fake *FakeMediaTrack) AddSubscriberCalls(stub func(types.Participant) error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = stub
}

func (fake *FakeMediaTrack) AddSubscriberArgsForCall(i int) types.Participant {
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	argsForCall := fake.addSubscriberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediaTrack) AddSubscriberReturns(result1 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	fake.addSubscriberReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeMediaTrack) AddSubscriberReturnsOnCall(i int, result1 error) {
	fake.addSubscriberMutex.Lock()
	defer fake.addSubscriberMutex.Unlock()
	fake.AddSubscriberStub = nil
	if fake.addSubscriberReturnsOnCall == nil {
		fake.addSubscriberReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addSubscriberReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeMediaTrack) GetQualityForDimension(arg1 uint32, arg2 uint32) livekit.VideoQuality {
	fake.getQualityForDimensionMutex.Lock()
	ret, specificReturn := fake.getQualityForDimensionReturnsOnCall[len(fake.getQualityForDimensionArgsForCall)]
	fake.getQualityForDimensionArgsForCall = append(fake.getQualityForDimensionArgsForCall, struct {
		arg1 uint32
		arg2 uint32
	}{arg1, arg2})
	stub := fake.GetQualityForDimensionStub
	fakeReturns := fake.getQualityForDimensionReturns
	fake.recordInvocation("GetQualityForDimension", []interface{}{arg1, arg2})
	fake.getQualityForDimensionMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) GetQualityForDimensionCallCount() int {
	fake.getQualityForDimensionMutex.RLock()
	defer fake.getQualityForDimensionMutex.RUnlock()
	return len(fake.getQualityForDimensionArgsForCall)
}

func (fake *FakeMediaTrack) GetQualityForDimensionCalls(stub func(uint32, uint32) livekit.VideoQuality) {
	fake.getQualityForDimensionMutex.Lock()
	defer fake.getQualityForDimensionMutex.Unlock()
	fake.GetQualityForDimensionStub = stub
}

func (fake *FakeMediaTrack) GetQualityForDimensionArgsForCall(i int) (uint32, uint32) {
	fake.getQualityForDimensionMutex.RLock()
	defer fake.getQualityForDimensionMutex.RUnlock()
	argsForCall := fake.getQualityForDimensionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMediaTrack) GetQualityForDimensionReturns(result1 livekit.VideoQuality) {
	fake.getQualityForDimensionMutex.Lock()
	defer fake.getQualityForDimensionMutex.Unlock()
	fake.GetQualityForDimensionStub = nil
	fake.getQualityForDimensionReturns = struct {
		result1 livekit.VideoQuality
	}{result1}
}

func (fake *FakeMediaTrack) GetQualityForDimensionReturnsOnCall(i int, result1 livekit.VideoQuality) {
	fake.getQualityForDimensionMutex.Lock()
	defer fake.getQualityForDimensionMutex.Unlock()
	fake.GetQualityForDimensionStub = nil
	if fake.getQualityForDimensionReturnsOnCall == nil {
		fake.getQualityForDimensionReturnsOnCall = make(map[int]struct {
			result1 livekit.VideoQuality
		})
	}
	fake.getQualityForDimensionReturnsOnCall[i] = struct {
		result1 livekit.VideoQuality
	}{result1}
}

func (fake *FakeMediaTrack) ID() string {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	stub := fake.IDStub
	fakeReturns := fake.iDReturns
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeMediaTrack) IDCalls(stub func() string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeMediaTrack) IDReturns(result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMediaTrack) IDReturnsOnCall(i int, result1 string) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeMediaTrack) IsMuted() bool {
	fake.isMutedMutex.Lock()
	ret, specificReturn := fake.isMutedReturnsOnCall[len(fake.isMutedArgsForCall)]
	fake.isMutedArgsForCall = append(fake.isMutedArgsForCall, struct {
	}{})
	stub := fake.IsMutedStub
	fakeReturns := fake.isMutedReturns
	fake.recordInvocation("IsMuted", []interface{}{})
	fake.isMutedMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) IsMutedCallCount() int {
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	return len(fake.isMutedArgsForCall)
}

func (fake *FakeMediaTrack) IsMutedCalls(stub func() bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = stub
}

func (fake *FakeMediaTrack) IsMutedReturns(result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	fake.isMutedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMediaTrack) IsMutedReturnsOnCall(i int, result1 bool) {
	fake.isMutedMutex.Lock()
	defer fake.isMutedMutex.Unlock()
	fake.IsMutedStub = nil
	if fake.isMutedReturnsOnCall == nil {
		fake.isMutedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isMutedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMediaTrack) IsSimulcast() bool {
	fake.isSimulcastMutex.Lock()
	ret, specificReturn := fake.isSimulcastReturnsOnCall[len(fake.isSimulcastArgsForCall)]
	fake.isSimulcastArgsForCall = append(fake.isSimulcastArgsForCall, struct {
	}{})
	stub := fake.IsSimulcastStub
	fakeReturns := fake.isSimulcastReturns
	fake.recordInvocation("IsSimulcast", []interface{}{})
	fake.isSimulcastMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) IsSimulcastCallCount() int {
	fake.isSimulcastMutex.RLock()
	defer fake.isSimulcastMutex.RUnlock()
	return len(fake.isSimulcastArgsForCall)
}

func (fake *FakeMediaTrack) IsSimulcastCalls(stub func() bool) {
	fake.isSimulcastMutex.Lock()
	defer fake.isSimulcastMutex.Unlock()
	fake.IsSimulcastStub = stub
}

func (fake *FakeMediaTrack) IsSimulcastReturns(result1 bool) {
	fake.isSimulcastMutex.Lock()
	defer fake.isSimulcastMutex.Unlock()
	fake.IsSimulcastStub = nil
	fake.isSimulcastReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMediaTrack) IsSimulcastReturnsOnCall(i int, result1 bool) {
	fake.isSimulcastMutex.Lock()
	defer fake.isSimulcastMutex.Unlock()
	fake.IsSimulcastStub = nil
	if fake.isSimulcastReturnsOnCall == nil {
		fake.isSimulcastReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isSimulcastReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMediaTrack) IsSubscriber(arg1 string) bool {
	fake.isSubscriberMutex.Lock()
	ret, specificReturn := fake.isSubscriberReturnsOnCall[len(fake.isSubscriberArgsForCall)]
	fake.isSubscriberArgsForCall = append(fake.isSubscriberArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.IsSubscriberStub
	fakeReturns := fake.isSubscriberReturns
	fake.recordInvocation("IsSubscriber", []interface{}{arg1})
	fake.isSubscriberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) IsSubscriberCallCount() int {
	fake.isSubscriberMutex.RLock()
	defer fake.isSubscriberMutex.RUnlock()
	return len(fake.isSubscriberArgsForCall)
}

func (fake *FakeMediaTrack) IsSubscriberCalls(stub func(string) bool) {
	fake.isSubscriberMutex.Lock()
	defer fake.isSubscriberMutex.Unlock()
	fake.IsSubscriberStub = stub
}

func (fake *FakeMediaTrack) IsSubscriberArgsForCall(i int) string {
	fake.isSubscriberMutex.RLock()
	defer fake.isSubscriberMutex.RUnlock()
	argsForCall := fake.isSubscriberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediaTrack) IsSubscriberReturns(result1 bool) {
	fake.isSubscriberMutex.Lock()
	defer fake.isSubscriberMutex.Unlock()
	fake.IsSubscriberStub = nil
	fake.isSubscriberReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMediaTrack) IsSubscriberReturnsOnCall(i int, result1 bool) {
	fake.isSubscriberMutex.Lock()
	defer fake.isSubscriberMutex.Unlock()
	fake.IsSubscriberStub = nil
	if fake.isSubscriberReturnsOnCall == nil {
		fake.isSubscriberReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isSubscriberReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeMediaTrack) Kind() livekit.TrackType {
	fake.kindMutex.Lock()
	ret, specificReturn := fake.kindReturnsOnCall[len(fake.kindArgsForCall)]
	fake.kindArgsForCall = append(fake.kindArgsForCall, struct {
	}{})
	stub := fake.KindStub
	fakeReturns := fake.kindReturns
	fake.recordInvocation("Kind", []interface{}{})
	fake.kindMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) KindCallCount() int {
	fake.kindMutex.RLock()
	defer fake.kindMutex.RUnlock()
	return len(fake.kindArgsForCall)
}

func (fake *FakeMediaTrack) KindCalls(stub func() livekit.TrackType) {
	fake.kindMutex.Lock()
	defer fake.kindMutex.Unlock()
	fake.KindStub = stub
}

func (fake *FakeMediaTrack) KindReturns(result1 livekit.TrackType) {
	fake.kindMutex.Lock()
	defer fake.kindMutex.Unlock()
	fake.KindStub = nil
	fake.kindReturns = struct {
		result1 livekit.TrackType
	}{result1}
}

func (fake *FakeMediaTrack) KindReturnsOnCall(i int, result1 livekit.TrackType) {
	fake.kindMutex.Lock()
	defer fake.kindMutex.Unlock()
	fake.KindStub = nil
	if fake.kindReturnsOnCall == nil {
		fake.kindReturnsOnCall = make(map[int]struct {
			result1 livekit.TrackType
		})
	}
	fake.kindReturnsOnCall[i] = struct {
		result1 livekit.TrackType
	}{result1}
}

func (fake *FakeMediaTrack) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	stub := fake.NameStub
	fakeReturns := fake.nameReturns
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeMediaTrack) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeMediaTrack) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeMediaTrack) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeMediaTrack) NotifySubscriberMaxQuality(arg1 string, arg2 livekit.VideoQuality) {
	fake.notifySubscriberMaxQualityMutex.Lock()
	fake.notifySubscriberMaxQualityArgsForCall = append(fake.notifySubscriberMaxQualityArgsForCall, struct {
		arg1 string
		arg2 livekit.VideoQuality
	}{arg1, arg2})
	stub := fake.NotifySubscriberMaxQualityStub
	fake.recordInvocation("NotifySubscriberMaxQuality", []interface{}{arg1, arg2})
	fake.notifySubscriberMaxQualityMutex.Unlock()
	if stub != nil {
		fake.NotifySubscriberMaxQualityStub(arg1, arg2)
	}
}

func (fake *FakeMediaTrack) NotifySubscriberMaxQualityCallCount() int {
	fake.notifySubscriberMaxQualityMutex.RLock()
	defer fake.notifySubscriberMaxQualityMutex.RUnlock()
	return len(fake.notifySubscriberMaxQualityArgsForCall)
}

func (fake *FakeMediaTrack) NotifySubscriberMaxQualityCalls(stub func(string, livekit.VideoQuality)) {
	fake.notifySubscriberMaxQualityMutex.Lock()
	defer fake.notifySubscriberMaxQualityMutex.Unlock()
	fake.NotifySubscriberMaxQualityStub = stub
}

func (fake *FakeMediaTrack) NotifySubscriberMaxQualityArgsForCall(i int) (string, livekit.VideoQuality) {
	fake.notifySubscriberMaxQualityMutex.RLock()
	defer fake.notifySubscriberMaxQualityMutex.RUnlock()
	argsForCall := fake.notifySubscriberMaxQualityArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeMediaTrack) NotifySubscriberMute(arg1 string) {
	fake.notifySubscriberMuteMutex.Lock()
	fake.notifySubscriberMuteArgsForCall = append(fake.notifySubscriberMuteArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.NotifySubscriberMuteStub
	fake.recordInvocation("NotifySubscriberMute", []interface{}{arg1})
	fake.notifySubscriberMuteMutex.Unlock()
	if stub != nil {
		fake.NotifySubscriberMuteStub(arg1)
	}
}

func (fake *FakeMediaTrack) NotifySubscriberMuteCallCount() int {
	fake.notifySubscriberMuteMutex.RLock()
	defer fake.notifySubscriberMuteMutex.RUnlock()
	return len(fake.notifySubscriberMuteArgsForCall)
}

func (fake *FakeMediaTrack) NotifySubscriberMuteCalls(stub func(string)) {
	fake.notifySubscriberMuteMutex.Lock()
	defer fake.notifySubscriberMuteMutex.Unlock()
	fake.NotifySubscriberMuteStub = stub
}

func (fake *FakeMediaTrack) NotifySubscriberMuteArgsForCall(i int) string {
	fake.notifySubscriberMuteMutex.RLock()
	defer fake.notifySubscriberMuteMutex.RUnlock()
	argsForCall := fake.notifySubscriberMuteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediaTrack) OnSubscribedMaxQualityChange(arg1 func(trackSid string, subscribedQualities []*livekit.SubscribedQuality) error) {
	fake.onSubscribedMaxQualityChangeMutex.Lock()
	fake.onSubscribedMaxQualityChangeArgsForCall = append(fake.onSubscribedMaxQualityChangeArgsForCall, struct {
		arg1 func(trackSid string, subscribedQualities []*livekit.SubscribedQuality) error
	}{arg1})
	stub := fake.OnSubscribedMaxQualityChangeStub
	fake.recordInvocation("OnSubscribedMaxQualityChange", []interface{}{arg1})
	fake.onSubscribedMaxQualityChangeMutex.Unlock()
	if stub != nil {
		fake.OnSubscribedMaxQualityChangeStub(arg1)
	}
}

func (fake *FakeMediaTrack) OnSubscribedMaxQualityChangeCallCount() int {
	fake.onSubscribedMaxQualityChangeMutex.RLock()
	defer fake.onSubscribedMaxQualityChangeMutex.RUnlock()
	return len(fake.onSubscribedMaxQualityChangeArgsForCall)
}

func (fake *FakeMediaTrack) OnSubscribedMaxQualityChangeCalls(stub func(func(trackSid string, subscribedQualities []*livekit.SubscribedQuality) error)) {
	fake.onSubscribedMaxQualityChangeMutex.Lock()
	defer fake.onSubscribedMaxQualityChangeMutex.Unlock()
	fake.OnSubscribedMaxQualityChangeStub = stub
}

func (fake *FakeMediaTrack) OnSubscribedMaxQualityChangeArgsForCall(i int) func(trackSid string, subscribedQualities []*livekit.SubscribedQuality) error {
	fake.onSubscribedMaxQualityChangeMutex.RLock()
	defer fake.onSubscribedMaxQualityChangeMutex.RUnlock()
	argsForCall := fake.onSubscribedMaxQualityChangeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediaTrack) RemoveAllSubscribers() {
	fake.removeAllSubscribersMutex.Lock()
	fake.removeAllSubscribersArgsForCall = append(fake.removeAllSubscribersArgsForCall, struct {
	}{})
	stub := fake.RemoveAllSubscribersStub
	fake.recordInvocation("RemoveAllSubscribers", []interface{}{})
	fake.removeAllSubscribersMutex.Unlock()
	if stub != nil {
		fake.RemoveAllSubscribersStub()
	}
}

func (fake *FakeMediaTrack) RemoveAllSubscribersCallCount() int {
	fake.removeAllSubscribersMutex.RLock()
	defer fake.removeAllSubscribersMutex.RUnlock()
	return len(fake.removeAllSubscribersArgsForCall)
}

func (fake *FakeMediaTrack) RemoveAllSubscribersCalls(stub func()) {
	fake.removeAllSubscribersMutex.Lock()
	defer fake.removeAllSubscribersMutex.Unlock()
	fake.RemoveAllSubscribersStub = stub
}

func (fake *FakeMediaTrack) RemoveSubscriber(arg1 string) {
	fake.removeSubscriberMutex.Lock()
	fake.removeSubscriberArgsForCall = append(fake.removeSubscriberArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemoveSubscriberStub
	fake.recordInvocation("RemoveSubscriber", []interface{}{arg1})
	fake.removeSubscriberMutex.Unlock()
	if stub != nil {
		fake.RemoveSubscriberStub(arg1)
	}
}

func (fake *FakeMediaTrack) RemoveSubscriberCallCount() int {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	return len(fake.removeSubscriberArgsForCall)
}

func (fake *FakeMediaTrack) RemoveSubscriberCalls(stub func(string)) {
	fake.removeSubscriberMutex.Lock()
	defer fake.removeSubscriberMutex.Unlock()
	fake.RemoveSubscriberStub = stub
}

func (fake *FakeMediaTrack) RemoveSubscriberArgsForCall(i int) string {
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	argsForCall := fake.removeSubscriberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediaTrack) SetMuted(arg1 bool) {
	fake.setMutedMutex.Lock()
	fake.setMutedArgsForCall = append(fake.setMutedArgsForCall, struct {
		arg1 bool
	}{arg1})
	stub := fake.SetMutedStub
	fake.recordInvocation("SetMuted", []interface{}{arg1})
	fake.setMutedMutex.Unlock()
	if stub != nil {
		fake.SetMutedStub(arg1)
	}
}

func (fake *FakeMediaTrack) SetMutedCallCount() int {
	fake.setMutedMutex.RLock()
	defer fake.setMutedMutex.RUnlock()
	return len(fake.setMutedArgsForCall)
}

func (fake *FakeMediaTrack) SetMutedCalls(stub func(bool)) {
	fake.setMutedMutex.Lock()
	defer fake.setMutedMutex.Unlock()
	fake.SetMutedStub = stub
}

func (fake *FakeMediaTrack) SetMutedArgsForCall(i int) bool {
	fake.setMutedMutex.RLock()
	defer fake.setMutedMutex.RUnlock()
	argsForCall := fake.setMutedArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediaTrack) Source() livekit.TrackSource {
	fake.sourceMutex.Lock()
	ret, specificReturn := fake.sourceReturnsOnCall[len(fake.sourceArgsForCall)]
	fake.sourceArgsForCall = append(fake.sourceArgsForCall, struct {
	}{})
	stub := fake.SourceStub
	fakeReturns := fake.sourceReturns
	fake.recordInvocation("Source", []interface{}{})
	fake.sourceMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeMediaTrack) SourceCallCount() int {
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	return len(fake.sourceArgsForCall)
}

func (fake *FakeMediaTrack) SourceCalls(stub func() livekit.TrackSource) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = stub
}

func (fake *FakeMediaTrack) SourceReturns(result1 livekit.TrackSource) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	fake.sourceReturns = struct {
		result1 livekit.TrackSource
	}{result1}
}

func (fake *FakeMediaTrack) SourceReturnsOnCall(i int, result1 livekit.TrackSource) {
	fake.sourceMutex.Lock()
	defer fake.sourceMutex.Unlock()
	fake.SourceStub = nil
	if fake.sourceReturnsOnCall == nil {
		fake.sourceReturnsOnCall = make(map[int]struct {
			result1 livekit.TrackSource
		})
	}
	fake.sourceReturnsOnCall[i] = struct {
		result1 livekit.TrackSource
	}{result1}
}

func (fake *FakeMediaTrack) UpdateVideoLayers(arg1 []*livekit.VideoLayer) {
	var arg1Copy []*livekit.VideoLayer
	if arg1 != nil {
		arg1Copy = make([]*livekit.VideoLayer, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.updateVideoLayersMutex.Lock()
	fake.updateVideoLayersArgsForCall = append(fake.updateVideoLayersArgsForCall, struct {
		arg1 []*livekit.VideoLayer
	}{arg1Copy})
	stub := fake.UpdateVideoLayersStub
	fake.recordInvocation("UpdateVideoLayers", []interface{}{arg1Copy})
	fake.updateVideoLayersMutex.Unlock()
	if stub != nil {
		fake.UpdateVideoLayersStub(arg1)
	}
}

func (fake *FakeMediaTrack) UpdateVideoLayersCallCount() int {
	fake.updateVideoLayersMutex.RLock()
	defer fake.updateVideoLayersMutex.RUnlock()
	return len(fake.updateVideoLayersArgsForCall)
}

func (fake *FakeMediaTrack) UpdateVideoLayersCalls(stub func([]*livekit.VideoLayer)) {
	fake.updateVideoLayersMutex.Lock()
	defer fake.updateVideoLayersMutex.Unlock()
	fake.UpdateVideoLayersStub = stub
}

func (fake *FakeMediaTrack) UpdateVideoLayersArgsForCall(i int) []*livekit.VideoLayer {
	fake.updateVideoLayersMutex.RLock()
	defer fake.updateVideoLayersMutex.RUnlock()
	argsForCall := fake.updateVideoLayersArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediaTrack) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addSubscriberMutex.RLock()
	defer fake.addSubscriberMutex.RUnlock()
	fake.getQualityForDimensionMutex.RLock()
	defer fake.getQualityForDimensionMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isMutedMutex.RLock()
	defer fake.isMutedMutex.RUnlock()
	fake.isSimulcastMutex.RLock()
	defer fake.isSimulcastMutex.RUnlock()
	fake.isSubscriberMutex.RLock()
	defer fake.isSubscriberMutex.RUnlock()
	fake.kindMutex.RLock()
	defer fake.kindMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.notifySubscriberMaxQualityMutex.RLock()
	defer fake.notifySubscriberMaxQualityMutex.RUnlock()
	fake.notifySubscriberMuteMutex.RLock()
	defer fake.notifySubscriberMuteMutex.RUnlock()
	fake.onSubscribedMaxQualityChangeMutex.RLock()
	defer fake.onSubscribedMaxQualityChangeMutex.RUnlock()
	fake.removeAllSubscribersMutex.RLock()
	defer fake.removeAllSubscribersMutex.RUnlock()
	fake.removeSubscriberMutex.RLock()
	defer fake.removeSubscriberMutex.RUnlock()
	fake.setMutedMutex.RLock()
	defer fake.setMutedMutex.RUnlock()
	fake.sourceMutex.RLock()
	defer fake.sourceMutex.RUnlock()
	fake.updateVideoLayersMutex.RLock()
	defer fake.updateVideoLayersMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMediaTrack) recordInvocation(key string, args []interface{}) {
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

var _ types.MediaTrack = new(FakeMediaTrack)
