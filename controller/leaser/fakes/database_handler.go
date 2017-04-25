// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"code.cloudfoundry.org/silk/controller"
)

type DatabaseHandler struct {
	AddEntryStub        func(controller.Lease) error
	addEntryMutex       sync.RWMutex
	addEntryArgsForCall []struct {
		arg1 controller.Lease
	}
	addEntryReturns struct {
		result1 error
	}
	addEntryReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteEntryStub        func(string) error
	deleteEntryMutex       sync.RWMutex
	deleteEntryArgsForCall []struct {
		arg1 string
	}
	deleteEntryReturns struct {
		result1 error
	}
	deleteEntryReturnsOnCall map[int]struct {
		result1 error
	}
	LeaseForUnderlayIPStub        func(string) (*controller.Lease, error)
	leaseForUnderlayIPMutex       sync.RWMutex
	leaseForUnderlayIPArgsForCall []struct {
		arg1 string
	}
	leaseForUnderlayIPReturns struct {
		result1 *controller.Lease
		result2 error
	}
	leaseForUnderlayIPReturnsOnCall map[int]struct {
		result1 *controller.Lease
		result2 error
	}
	LastRenewedAtForUnderlayIPStub        func(string) (int64, error)
	lastRenewedAtForUnderlayIPMutex       sync.RWMutex
	lastRenewedAtForUnderlayIPArgsForCall []struct {
		arg1 string
	}
	lastRenewedAtForUnderlayIPReturns struct {
		result1 int64
		result2 error
	}
	lastRenewedAtForUnderlayIPReturnsOnCall map[int]struct {
		result1 int64
		result2 error
	}
	RenewLeaseForUnderlayIPStub        func(string) error
	renewLeaseForUnderlayIPMutex       sync.RWMutex
	renewLeaseForUnderlayIPArgsForCall []struct {
		arg1 string
	}
	renewLeaseForUnderlayIPReturns struct {
		result1 error
	}
	renewLeaseForUnderlayIPReturnsOnCall map[int]struct {
		result1 error
	}
	AllStub        func() ([]controller.Lease, error)
	allMutex       sync.RWMutex
	allArgsForCall []struct{}
	allReturns     struct {
		result1 []controller.Lease
		result2 error
	}
	allReturnsOnCall map[int]struct {
		result1 []controller.Lease
		result2 error
	}
	OldestExpiredStub        func(int) (*controller.Lease, error)
	oldestExpiredMutex       sync.RWMutex
	oldestExpiredArgsForCall []struct {
		arg1 int
	}
	oldestExpiredReturns struct {
		result1 *controller.Lease
		result2 error
	}
	oldestExpiredReturnsOnCall map[int]struct {
		result1 *controller.Lease
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DatabaseHandler) AddEntry(arg1 controller.Lease) error {
	fake.addEntryMutex.Lock()
	ret, specificReturn := fake.addEntryReturnsOnCall[len(fake.addEntryArgsForCall)]
	fake.addEntryArgsForCall = append(fake.addEntryArgsForCall, struct {
		arg1 controller.Lease
	}{arg1})
	fake.recordInvocation("AddEntry", []interface{}{arg1})
	fake.addEntryMutex.Unlock()
	if fake.AddEntryStub != nil {
		return fake.AddEntryStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addEntryReturns.result1
}

func (fake *DatabaseHandler) AddEntryCallCount() int {
	fake.addEntryMutex.RLock()
	defer fake.addEntryMutex.RUnlock()
	return len(fake.addEntryArgsForCall)
}

func (fake *DatabaseHandler) AddEntryArgsForCall(i int) controller.Lease {
	fake.addEntryMutex.RLock()
	defer fake.addEntryMutex.RUnlock()
	return fake.addEntryArgsForCall[i].arg1
}

func (fake *DatabaseHandler) AddEntryReturns(result1 error) {
	fake.AddEntryStub = nil
	fake.addEntryReturns = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) AddEntryReturnsOnCall(i int, result1 error) {
	fake.AddEntryStub = nil
	if fake.addEntryReturnsOnCall == nil {
		fake.addEntryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addEntryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) DeleteEntry(arg1 string) error {
	fake.deleteEntryMutex.Lock()
	ret, specificReturn := fake.deleteEntryReturnsOnCall[len(fake.deleteEntryArgsForCall)]
	fake.deleteEntryArgsForCall = append(fake.deleteEntryArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("DeleteEntry", []interface{}{arg1})
	fake.deleteEntryMutex.Unlock()
	if fake.DeleteEntryStub != nil {
		return fake.DeleteEntryStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deleteEntryReturns.result1
}

func (fake *DatabaseHandler) DeleteEntryCallCount() int {
	fake.deleteEntryMutex.RLock()
	defer fake.deleteEntryMutex.RUnlock()
	return len(fake.deleteEntryArgsForCall)
}

func (fake *DatabaseHandler) DeleteEntryArgsForCall(i int) string {
	fake.deleteEntryMutex.RLock()
	defer fake.deleteEntryMutex.RUnlock()
	return fake.deleteEntryArgsForCall[i].arg1
}

func (fake *DatabaseHandler) DeleteEntryReturns(result1 error) {
	fake.DeleteEntryStub = nil
	fake.deleteEntryReturns = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) DeleteEntryReturnsOnCall(i int, result1 error) {
	fake.DeleteEntryStub = nil
	if fake.deleteEntryReturnsOnCall == nil {
		fake.deleteEntryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteEntryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) LeaseForUnderlayIP(arg1 string) (*controller.Lease, error) {
	fake.leaseForUnderlayIPMutex.Lock()
	ret, specificReturn := fake.leaseForUnderlayIPReturnsOnCall[len(fake.leaseForUnderlayIPArgsForCall)]
	fake.leaseForUnderlayIPArgsForCall = append(fake.leaseForUnderlayIPArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LeaseForUnderlayIP", []interface{}{arg1})
	fake.leaseForUnderlayIPMutex.Unlock()
	if fake.LeaseForUnderlayIPStub != nil {
		return fake.LeaseForUnderlayIPStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.leaseForUnderlayIPReturns.result1, fake.leaseForUnderlayIPReturns.result2
}

func (fake *DatabaseHandler) LeaseForUnderlayIPCallCount() int {
	fake.leaseForUnderlayIPMutex.RLock()
	defer fake.leaseForUnderlayIPMutex.RUnlock()
	return len(fake.leaseForUnderlayIPArgsForCall)
}

func (fake *DatabaseHandler) LeaseForUnderlayIPArgsForCall(i int) string {
	fake.leaseForUnderlayIPMutex.RLock()
	defer fake.leaseForUnderlayIPMutex.RUnlock()
	return fake.leaseForUnderlayIPArgsForCall[i].arg1
}

func (fake *DatabaseHandler) LeaseForUnderlayIPReturns(result1 *controller.Lease, result2 error) {
	fake.LeaseForUnderlayIPStub = nil
	fake.leaseForUnderlayIPReturns = struct {
		result1 *controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) LeaseForUnderlayIPReturnsOnCall(i int, result1 *controller.Lease, result2 error) {
	fake.LeaseForUnderlayIPStub = nil
	if fake.leaseForUnderlayIPReturnsOnCall == nil {
		fake.leaseForUnderlayIPReturnsOnCall = make(map[int]struct {
			result1 *controller.Lease
			result2 error
		})
	}
	fake.leaseForUnderlayIPReturnsOnCall[i] = struct {
		result1 *controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) LastRenewedAtForUnderlayIP(arg1 string) (int64, error) {
	fake.lastRenewedAtForUnderlayIPMutex.Lock()
	ret, specificReturn := fake.lastRenewedAtForUnderlayIPReturnsOnCall[len(fake.lastRenewedAtForUnderlayIPArgsForCall)]
	fake.lastRenewedAtForUnderlayIPArgsForCall = append(fake.lastRenewedAtForUnderlayIPArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("LastRenewedAtForUnderlayIP", []interface{}{arg1})
	fake.lastRenewedAtForUnderlayIPMutex.Unlock()
	if fake.LastRenewedAtForUnderlayIPStub != nil {
		return fake.LastRenewedAtForUnderlayIPStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lastRenewedAtForUnderlayIPReturns.result1, fake.lastRenewedAtForUnderlayIPReturns.result2
}

func (fake *DatabaseHandler) LastRenewedAtForUnderlayIPCallCount() int {
	fake.lastRenewedAtForUnderlayIPMutex.RLock()
	defer fake.lastRenewedAtForUnderlayIPMutex.RUnlock()
	return len(fake.lastRenewedAtForUnderlayIPArgsForCall)
}

func (fake *DatabaseHandler) LastRenewedAtForUnderlayIPArgsForCall(i int) string {
	fake.lastRenewedAtForUnderlayIPMutex.RLock()
	defer fake.lastRenewedAtForUnderlayIPMutex.RUnlock()
	return fake.lastRenewedAtForUnderlayIPArgsForCall[i].arg1
}

func (fake *DatabaseHandler) LastRenewedAtForUnderlayIPReturns(result1 int64, result2 error) {
	fake.LastRenewedAtForUnderlayIPStub = nil
	fake.lastRenewedAtForUnderlayIPReturns = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) LastRenewedAtForUnderlayIPReturnsOnCall(i int, result1 int64, result2 error) {
	fake.LastRenewedAtForUnderlayIPStub = nil
	if fake.lastRenewedAtForUnderlayIPReturnsOnCall == nil {
		fake.lastRenewedAtForUnderlayIPReturnsOnCall = make(map[int]struct {
			result1 int64
			result2 error
		})
	}
	fake.lastRenewedAtForUnderlayIPReturnsOnCall[i] = struct {
		result1 int64
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) RenewLeaseForUnderlayIP(arg1 string) error {
	fake.renewLeaseForUnderlayIPMutex.Lock()
	ret, specificReturn := fake.renewLeaseForUnderlayIPReturnsOnCall[len(fake.renewLeaseForUnderlayIPArgsForCall)]
	fake.renewLeaseForUnderlayIPArgsForCall = append(fake.renewLeaseForUnderlayIPArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("RenewLeaseForUnderlayIP", []interface{}{arg1})
	fake.renewLeaseForUnderlayIPMutex.Unlock()
	if fake.RenewLeaseForUnderlayIPStub != nil {
		return fake.RenewLeaseForUnderlayIPStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.renewLeaseForUnderlayIPReturns.result1
}

func (fake *DatabaseHandler) RenewLeaseForUnderlayIPCallCount() int {
	fake.renewLeaseForUnderlayIPMutex.RLock()
	defer fake.renewLeaseForUnderlayIPMutex.RUnlock()
	return len(fake.renewLeaseForUnderlayIPArgsForCall)
}

func (fake *DatabaseHandler) RenewLeaseForUnderlayIPArgsForCall(i int) string {
	fake.renewLeaseForUnderlayIPMutex.RLock()
	defer fake.renewLeaseForUnderlayIPMutex.RUnlock()
	return fake.renewLeaseForUnderlayIPArgsForCall[i].arg1
}

func (fake *DatabaseHandler) RenewLeaseForUnderlayIPReturns(result1 error) {
	fake.RenewLeaseForUnderlayIPStub = nil
	fake.renewLeaseForUnderlayIPReturns = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) RenewLeaseForUnderlayIPReturnsOnCall(i int, result1 error) {
	fake.RenewLeaseForUnderlayIPStub = nil
	if fake.renewLeaseForUnderlayIPReturnsOnCall == nil {
		fake.renewLeaseForUnderlayIPReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.renewLeaseForUnderlayIPReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *DatabaseHandler) All() ([]controller.Lease, error) {
	fake.allMutex.Lock()
	ret, specificReturn := fake.allReturnsOnCall[len(fake.allArgsForCall)]
	fake.allArgsForCall = append(fake.allArgsForCall, struct{}{})
	fake.recordInvocation("All", []interface{}{})
	fake.allMutex.Unlock()
	if fake.AllStub != nil {
		return fake.AllStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.allReturns.result1, fake.allReturns.result2
}

func (fake *DatabaseHandler) AllCallCount() int {
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	return len(fake.allArgsForCall)
}

func (fake *DatabaseHandler) AllReturns(result1 []controller.Lease, result2 error) {
	fake.AllStub = nil
	fake.allReturns = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) AllReturnsOnCall(i int, result1 []controller.Lease, result2 error) {
	fake.AllStub = nil
	if fake.allReturnsOnCall == nil {
		fake.allReturnsOnCall = make(map[int]struct {
			result1 []controller.Lease
			result2 error
		})
	}
	fake.allReturnsOnCall[i] = struct {
		result1 []controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) OldestExpired(arg1 int) (*controller.Lease, error) {
	fake.oldestExpiredMutex.Lock()
	ret, specificReturn := fake.oldestExpiredReturnsOnCall[len(fake.oldestExpiredArgsForCall)]
	fake.oldestExpiredArgsForCall = append(fake.oldestExpiredArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("OldestExpired", []interface{}{arg1})
	fake.oldestExpiredMutex.Unlock()
	if fake.OldestExpiredStub != nil {
		return fake.OldestExpiredStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.oldestExpiredReturns.result1, fake.oldestExpiredReturns.result2
}

func (fake *DatabaseHandler) OldestExpiredCallCount() int {
	fake.oldestExpiredMutex.RLock()
	defer fake.oldestExpiredMutex.RUnlock()
	return len(fake.oldestExpiredArgsForCall)
}

func (fake *DatabaseHandler) OldestExpiredArgsForCall(i int) int {
	fake.oldestExpiredMutex.RLock()
	defer fake.oldestExpiredMutex.RUnlock()
	return fake.oldestExpiredArgsForCall[i].arg1
}

func (fake *DatabaseHandler) OldestExpiredReturns(result1 *controller.Lease, result2 error) {
	fake.OldestExpiredStub = nil
	fake.oldestExpiredReturns = struct {
		result1 *controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) OldestExpiredReturnsOnCall(i int, result1 *controller.Lease, result2 error) {
	fake.OldestExpiredStub = nil
	if fake.oldestExpiredReturnsOnCall == nil {
		fake.oldestExpiredReturnsOnCall = make(map[int]struct {
			result1 *controller.Lease
			result2 error
		})
	}
	fake.oldestExpiredReturnsOnCall[i] = struct {
		result1 *controller.Lease
		result2 error
	}{result1, result2}
}

func (fake *DatabaseHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addEntryMutex.RLock()
	defer fake.addEntryMutex.RUnlock()
	fake.deleteEntryMutex.RLock()
	defer fake.deleteEntryMutex.RUnlock()
	fake.leaseForUnderlayIPMutex.RLock()
	defer fake.leaseForUnderlayIPMutex.RUnlock()
	fake.lastRenewedAtForUnderlayIPMutex.RLock()
	defer fake.lastRenewedAtForUnderlayIPMutex.RUnlock()
	fake.renewLeaseForUnderlayIPMutex.RLock()
	defer fake.renewLeaseForUnderlayIPMutex.RUnlock()
	fake.allMutex.RLock()
	defer fake.allMutex.RUnlock()
	fake.oldestExpiredMutex.RLock()
	defer fake.oldestExpiredMutex.RUnlock()
	return fake.invocations
}

func (fake *DatabaseHandler) recordInvocation(key string, args []interface{}) {
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
