package go_promise

import (
	"fmt"
	"strconv"
	"sync"
)

type GoPromise struct {
	wg           *sync.WaitGroup
	state        PromiseState
	callbacks    []func(any) any
	errorHandler func(string)
}

func New(callback func(func(any), func(string))) *GoPromise {
	fmt.Println("A")
	p := &(GoPromise{})
	p.wg = &sync.WaitGroup{}
	p.wg.Add(1)
	p.state = PENDING
	p.callbacks = [](func(any) any){}
	callback(p.Resolve, p.Reject)
	return p
}

func (p *GoPromise) Resolve(data any) {
	p.state = RESOLVED
	fmt.Println("D")
	p.runCallbacks(data)
	p.wg.Done()
}

func (p *GoPromise) Reject(errMsg string) {
	p.state = REJECTED
	p.errorHandler(errMsg)
	p.wg.Done()
}

func (p *GoPromise) Then(callback func(any) any) *GoPromise {
	p.callbacks = append(p.callbacks, callback)
	fmt.Println("C")
	fmt.Println("callbacks len: " + strconv.Itoa(len(p.callbacks)))
	fmt.Println(p)
	return p
}

func (p *GoPromise) Catch(handler func(string)) {
	p.errorHandler = handler
}

func (p *GoPromise) runCallbacks(initialData any) {
	fmt.Println("E")
	fmt.Println("callbacks len: " + strconv.Itoa(len(p.callbacks)))
	fmt.Println(p)
	data := initialData
	for _, callback := range p.callbacks {
		data = callback(data)
	}
}

func (p *GoPromise) Wait() {
	p.wg.Wait()
}
