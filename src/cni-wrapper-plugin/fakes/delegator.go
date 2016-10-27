// This file was generated by counterfeiter
package fakes

import (
	"cni-wrapper-plugin/lib"
	"sync"

	"github.com/containernetworking/cni/pkg/types"
)

type Delegator struct {
	DelegateAddStub        func(delegatePlugin string, netconf []byte) (*types.Result, error)
	delegateAddMutex       sync.RWMutex
	delegateAddArgsForCall []struct {
		delegatePlugin string
		netconf        []byte
	}
	delegateAddReturns struct {
		result1 *types.Result
		result2 error
	}
	DelegateDelStub        func(delegatePlugin string, netconf []byte) error
	delegateDelMutex       sync.RWMutex
	delegateDelArgsForCall []struct {
		delegatePlugin string
		netconf        []byte
	}
	delegateDelReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Delegator) DelegateAdd(delegatePlugin string, netconf []byte) (*types.Result, error) {
	var netconfCopy []byte
	if netconf != nil {
		netconfCopy = make([]byte, len(netconf))
		copy(netconfCopy, netconf)
	}
	fake.delegateAddMutex.Lock()
	fake.delegateAddArgsForCall = append(fake.delegateAddArgsForCall, struct {
		delegatePlugin string
		netconf        []byte
	}{delegatePlugin, netconfCopy})
	fake.recordInvocation("DelegateAdd", []interface{}{delegatePlugin, netconfCopy})
	fake.delegateAddMutex.Unlock()
	if fake.DelegateAddStub != nil {
		return fake.DelegateAddStub(delegatePlugin, netconf)
	} else {
		return fake.delegateAddReturns.result1, fake.delegateAddReturns.result2
	}
}

func (fake *Delegator) DelegateAddCallCount() int {
	fake.delegateAddMutex.RLock()
	defer fake.delegateAddMutex.RUnlock()
	return len(fake.delegateAddArgsForCall)
}

func (fake *Delegator) DelegateAddArgsForCall(i int) (string, []byte) {
	fake.delegateAddMutex.RLock()
	defer fake.delegateAddMutex.RUnlock()
	return fake.delegateAddArgsForCall[i].delegatePlugin, fake.delegateAddArgsForCall[i].netconf
}

func (fake *Delegator) DelegateAddReturns(result1 *types.Result, result2 error) {
	fake.DelegateAddStub = nil
	fake.delegateAddReturns = struct {
		result1 *types.Result
		result2 error
	}{result1, result2}
}

func (fake *Delegator) DelegateDel(delegatePlugin string, netconf []byte) error {
	var netconfCopy []byte
	if netconf != nil {
		netconfCopy = make([]byte, len(netconf))
		copy(netconfCopy, netconf)
	}
	fake.delegateDelMutex.Lock()
	fake.delegateDelArgsForCall = append(fake.delegateDelArgsForCall, struct {
		delegatePlugin string
		netconf        []byte
	}{delegatePlugin, netconfCopy})
	fake.recordInvocation("DelegateDel", []interface{}{delegatePlugin, netconfCopy})
	fake.delegateDelMutex.Unlock()
	if fake.DelegateDelStub != nil {
		return fake.DelegateDelStub(delegatePlugin, netconf)
	} else {
		return fake.delegateDelReturns.result1
	}
}

func (fake *Delegator) DelegateDelCallCount() int {
	fake.delegateDelMutex.RLock()
	defer fake.delegateDelMutex.RUnlock()
	return len(fake.delegateDelArgsForCall)
}

func (fake *Delegator) DelegateDelArgsForCall(i int) (string, []byte) {
	fake.delegateDelMutex.RLock()
	defer fake.delegateDelMutex.RUnlock()
	return fake.delegateDelArgsForCall[i].delegatePlugin, fake.delegateDelArgsForCall[i].netconf
}

func (fake *Delegator) DelegateDelReturns(result1 error) {
	fake.DelegateDelStub = nil
	fake.delegateDelReturns = struct {
		result1 error
	}{result1}
}

func (fake *Delegator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.delegateAddMutex.RLock()
	defer fake.delegateAddMutex.RUnlock()
	fake.delegateDelMutex.RLock()
	defer fake.delegateDelMutex.RUnlock()
	return fake.invocations
}

func (fake *Delegator) recordInvocation(key string, args []interface{}) {
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

var _ lib.Delegator = new(Delegator)
