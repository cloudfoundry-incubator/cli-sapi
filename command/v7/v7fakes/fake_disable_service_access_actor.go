// Code generated by counterfeiter. DO NOT EDIT.
package v7fakes

import (
	"sync"

	"code.cloudfoundry.org/cli/actor/v7action"
	v7 "code.cloudfoundry.org/cli/command/v7"
)

type FakeDisableServiceAccessActor struct {
	DisableServiceAccessStub        func(string, string, string, string) (v7action.SkippedPlans, v7action.Warnings, error)
	disableServiceAccessMutex       sync.RWMutex
	disableServiceAccessArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	disableServiceAccessReturns struct {
		result1 v7action.SkippedPlans
		result2 v7action.Warnings
		result3 error
	}
	disableServiceAccessReturnsOnCall map[int]struct {
		result1 v7action.SkippedPlans
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDisableServiceAccessActor) DisableServiceAccess(arg1 string, arg2 string, arg3 string, arg4 string) (v7action.SkippedPlans, v7action.Warnings, error) {
	fake.disableServiceAccessMutex.Lock()
	ret, specificReturn := fake.disableServiceAccessReturnsOnCall[len(fake.disableServiceAccessArgsForCall)]
	fake.disableServiceAccessArgsForCall = append(fake.disableServiceAccessArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("DisableServiceAccess", []interface{}{arg1, arg2, arg3, arg4})
	fake.disableServiceAccessMutex.Unlock()
	if fake.DisableServiceAccessStub != nil {
		return fake.DisableServiceAccessStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.disableServiceAccessReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeDisableServiceAccessActor) DisableServiceAccessCallCount() int {
	fake.disableServiceAccessMutex.RLock()
	defer fake.disableServiceAccessMutex.RUnlock()
	return len(fake.disableServiceAccessArgsForCall)
}

func (fake *FakeDisableServiceAccessActor) DisableServiceAccessCalls(stub func(string, string, string, string) (v7action.SkippedPlans, v7action.Warnings, error)) {
	fake.disableServiceAccessMutex.Lock()
	defer fake.disableServiceAccessMutex.Unlock()
	fake.DisableServiceAccessStub = stub
}

func (fake *FakeDisableServiceAccessActor) DisableServiceAccessArgsForCall(i int) (string, string, string, string) {
	fake.disableServiceAccessMutex.RLock()
	defer fake.disableServiceAccessMutex.RUnlock()
	argsForCall := fake.disableServiceAccessArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeDisableServiceAccessActor) DisableServiceAccessReturns(result1 v7action.SkippedPlans, result2 v7action.Warnings, result3 error) {
	fake.disableServiceAccessMutex.Lock()
	defer fake.disableServiceAccessMutex.Unlock()
	fake.DisableServiceAccessStub = nil
	fake.disableServiceAccessReturns = struct {
		result1 v7action.SkippedPlans
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDisableServiceAccessActor) DisableServiceAccessReturnsOnCall(i int, result1 v7action.SkippedPlans, result2 v7action.Warnings, result3 error) {
	fake.disableServiceAccessMutex.Lock()
	defer fake.disableServiceAccessMutex.Unlock()
	fake.DisableServiceAccessStub = nil
	if fake.disableServiceAccessReturnsOnCall == nil {
		fake.disableServiceAccessReturnsOnCall = make(map[int]struct {
			result1 v7action.SkippedPlans
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.disableServiceAccessReturnsOnCall[i] = struct {
		result1 v7action.SkippedPlans
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeDisableServiceAccessActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.disableServiceAccessMutex.RLock()
	defer fake.disableServiceAccessMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDisableServiceAccessActor) recordInvocation(key string, args []interface{}) {
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

var _ v7.DisableServiceAccessActor = new(FakeDisableServiceAccessActor)
