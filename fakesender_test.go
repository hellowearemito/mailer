// Code generated by counterfeiter. DO NOT EDIT.
package mailer_test

import (
	"sync"

	"github.com/nkovacs/gophermail"
)

type FakeSender struct {
	SendMailStub        func(msg *gophermail.Message) error
	sendMailMutex       sync.RWMutex
	sendMailArgsForCall []struct {
		msg *gophermail.Message
	}
	sendMailReturns *struct {
		result1 error
	}
	sendMailReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSender) SendMail(msg *gophermail.Message) error {
	fake.sendMailMutex.Lock()
	ret, specificReturn := fake.sendMailReturnsOnCall[len(fake.sendMailArgsForCall)]
	fake.sendMailArgsForCall = append(fake.sendMailArgsForCall, struct {
		msg *gophermail.Message
	}{msg})
	fake.recordInvocation("SendMail", []interface{}{msg})
	fake.sendMailMutex.Unlock()
	if fake.SendMailStub != nil {
		return fake.SendMailStub(msg)
	}
	if specificReturn {
		return ret.result1
	}
	if fake.sendMailReturns == nil {
		panic("Unexpected method call: Sender.SendMail()")
	}
	return fake.sendMailReturns.result1
}

func (fake *FakeSender) SendMailCallCount() int {
	fake.sendMailMutex.RLock()
	defer fake.sendMailMutex.RUnlock()
	return len(fake.sendMailArgsForCall)
}

func (fake *FakeSender) SendMailArgsForCall(i int) *gophermail.Message {
	fake.sendMailMutex.RLock()
	defer fake.sendMailMutex.RUnlock()
	return fake.sendMailArgsForCall[i].msg
}

func (fake *FakeSender) SendMailReturns(result1 error) {
	fake.SendMailStub = nil
	fake.sendMailReturns = &struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) SendMailReturnsOnCall(i int, result1 error) {
	fake.SendMailStub = nil
	if fake.sendMailReturnsOnCall == nil {
		fake.sendMailReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.sendMailReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.sendMailMutex.RLock()
	defer fake.sendMailMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSender) recordInvocation(key string, args []interface{}) {
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

var _ gophermail.Sender = new(FakeSender)
