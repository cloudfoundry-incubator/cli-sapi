// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeSpaceActor struct {
	GetEffectiveIsolationSegmentBySpaceStub        func(string, string) (v7action.IsolationSegment, v7action.Warnings, error)
	getEffectiveIsolationSegmentBySpaceMutex       sync.RWMutex
	getEffectiveIsolationSegmentBySpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getEffectiveIsolationSegmentBySpaceReturns struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}
	getEffectiveIsolationSegmentBySpaceReturnsOnCall map[int]struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}
	GetSpaceByNameAndOrganizationStub        func(string, string) (v7action.Space, v7action.Warnings, error)
	getSpaceByNameAndOrganizationMutex       sync.RWMutex
	getSpaceByNameAndOrganizationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceByNameAndOrganizationReturns struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	getSpaceByNameAndOrganizationReturnsOnCall map[int]struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}
	GetSpaceSummaryByNameAndOrganizationStub        func(string, string) (v7action.SpaceSummary, v7action.Warnings, error)
	getSpaceSummaryByNameAndOrganizationMutex       sync.RWMutex
	getSpaceSummaryByNameAndOrganizationArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceSummaryByNameAndOrganizationReturns struct {
		result1 v7action.SpaceSummary
		result2 v7action.Warnings
		result3 error
	}
	getSpaceSummaryByNameAndOrganizationReturnsOnCall map[int]struct {
		result1 v7action.SpaceSummary
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSpaceActor) GetEffectiveIsolationSegmentBySpace(arg1 string, arg2 string) (v7action.IsolationSegment, v7action.Warnings, error) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	ret, specificReturn := fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall[len(fake.getEffectiveIsolationSegmentBySpaceArgsForCall)]
	fake.getEffectiveIsolationSegmentBySpaceArgsForCall = append(fake.getEffectiveIsolationSegmentBySpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetEffectiveIsolationSegmentBySpace", []interface{}{arg1, arg2})
	fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	if fake.GetEffectiveIsolationSegmentBySpaceStub != nil {
		return fake.GetEffectiveIsolationSegmentBySpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getEffectiveIsolationSegmentBySpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceActor) GetEffectiveIsolationSegmentBySpaceCallCount() int {
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	return len(fake.getEffectiveIsolationSegmentBySpaceArgsForCall)
}

func (fake *FakeSpaceActor) GetEffectiveIsolationSegmentBySpaceCalls(stub func(string, string) (v7action.IsolationSegment, v7action.Warnings, error)) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	fake.GetEffectiveIsolationSegmentBySpaceStub = stub
}

func (fake *FakeSpaceActor) GetEffectiveIsolationSegmentBySpaceArgsForCall(i int) (string, string) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	argsForCall := fake.getEffectiveIsolationSegmentBySpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpaceActor) GetEffectiveIsolationSegmentBySpaceReturns(result1 v7action.IsolationSegment, result2 v7action.Warnings, result3 error) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	fake.GetEffectiveIsolationSegmentBySpaceStub = nil
	fake.getEffectiveIsolationSegmentBySpaceReturns = struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetEffectiveIsolationSegmentBySpaceReturnsOnCall(i int, result1 v7action.IsolationSegment, result2 v7action.Warnings, result3 error) {
	fake.getEffectiveIsolationSegmentBySpaceMutex.Lock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.Unlock()
	fake.GetEffectiveIsolationSegmentBySpaceStub = nil
	if fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall == nil {
		fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.IsolationSegment
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getEffectiveIsolationSegmentBySpaceReturnsOnCall[i] = struct {
		result1 v7action.IsolationSegment
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetSpaceByNameAndOrganization(arg1 string, arg2 string) (v7action.Space, v7action.Warnings, error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	ret, specificReturn := fake.getSpaceByNameAndOrganizationReturnsOnCall[len(fake.getSpaceByNameAndOrganizationArgsForCall)]
	fake.getSpaceByNameAndOrganizationArgsForCall = append(fake.getSpaceByNameAndOrganizationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceByNameAndOrganization", []interface{}{arg1, arg2})
	fake.getSpaceByNameAndOrganizationMutex.Unlock()
	if fake.GetSpaceByNameAndOrganizationStub != nil {
		return fake.GetSpaceByNameAndOrganizationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceByNameAndOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceActor) GetSpaceByNameAndOrganizationCallCount() int {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceByNameAndOrganizationArgsForCall)
}

func (fake *FakeSpaceActor) GetSpaceByNameAndOrganizationCalls(stub func(string, string) (v7action.Space, v7action.Warnings, error)) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = stub
}

func (fake *FakeSpaceActor) GetSpaceByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	argsForCall := fake.getSpaceByNameAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpaceActor) GetSpaceByNameAndOrganizationReturns(result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	fake.getSpaceByNameAndOrganizationReturns = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetSpaceByNameAndOrganizationReturnsOnCall(i int, result1 v7action.Space, result2 v7action.Warnings, result3 error) {
	fake.getSpaceByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceByNameAndOrganizationStub = nil
	if fake.getSpaceByNameAndOrganizationReturnsOnCall == nil {
		fake.getSpaceByNameAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v7action.Space
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceByNameAndOrganizationReturnsOnCall[i] = struct {
		result1 v7action.Space
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetSpaceSummaryByNameAndOrganization(arg1 string, arg2 string) (v7action.SpaceSummary, v7action.Warnings, error) {
	fake.getSpaceSummaryByNameAndOrganizationMutex.Lock()
	ret, specificReturn := fake.getSpaceSummaryByNameAndOrganizationReturnsOnCall[len(fake.getSpaceSummaryByNameAndOrganizationArgsForCall)]
	fake.getSpaceSummaryByNameAndOrganizationArgsForCall = append(fake.getSpaceSummaryByNameAndOrganizationArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetSpaceSummaryByNameAndOrganization", []interface{}{arg1, arg2})
	fake.getSpaceSummaryByNameAndOrganizationMutex.Unlock()
	if fake.GetSpaceSummaryByNameAndOrganizationStub != nil {
		return fake.GetSpaceSummaryByNameAndOrganizationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSpaceSummaryByNameAndOrganizationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeSpaceActor) GetSpaceSummaryByNameAndOrganizationCallCount() int {
	fake.getSpaceSummaryByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceSummaryByNameAndOrganizationMutex.RUnlock()
	return len(fake.getSpaceSummaryByNameAndOrganizationArgsForCall)
}

func (fake *FakeSpaceActor) GetSpaceSummaryByNameAndOrganizationCalls(stub func(string, string) (v7action.SpaceSummary, v7action.Warnings, error)) {
	fake.getSpaceSummaryByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceSummaryByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceSummaryByNameAndOrganizationStub = stub
}

func (fake *FakeSpaceActor) GetSpaceSummaryByNameAndOrganizationArgsForCall(i int) (string, string) {
	fake.getSpaceSummaryByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceSummaryByNameAndOrganizationMutex.RUnlock()
	argsForCall := fake.getSpaceSummaryByNameAndOrganizationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeSpaceActor) GetSpaceSummaryByNameAndOrganizationReturns(result1 v7action.SpaceSummary, result2 v7action.Warnings, result3 error) {
	fake.getSpaceSummaryByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceSummaryByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceSummaryByNameAndOrganizationStub = nil
	fake.getSpaceSummaryByNameAndOrganizationReturns = struct {
		result1 v7action.SpaceSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) GetSpaceSummaryByNameAndOrganizationReturnsOnCall(i int, result1 v7action.SpaceSummary, result2 v7action.Warnings, result3 error) {
	fake.getSpaceSummaryByNameAndOrganizationMutex.Lock()
	defer fake.getSpaceSummaryByNameAndOrganizationMutex.Unlock()
	fake.GetSpaceSummaryByNameAndOrganizationStub = nil
	if fake.getSpaceSummaryByNameAndOrganizationReturnsOnCall == nil {
		fake.getSpaceSummaryByNameAndOrganizationReturnsOnCall = make(map[int]struct {
			result1 v7action.SpaceSummary
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getSpaceSummaryByNameAndOrganizationReturnsOnCall[i] = struct {
		result1 v7action.SpaceSummary
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeSpaceActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getEffectiveIsolationSegmentBySpaceMutex.RLock()
	defer fake.getEffectiveIsolationSegmentBySpaceMutex.RUnlock()
	fake.getSpaceByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceByNameAndOrganizationMutex.RUnlock()
	fake.getSpaceSummaryByNameAndOrganizationMutex.RLock()
	defer fake.getSpaceSummaryByNameAndOrganizationMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSpaceActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.SpaceActor = new(FakeSpaceActor)
