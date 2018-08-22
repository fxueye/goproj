package server

import (
	"sync"
)

type IService interface {
	Start() error
	Close()
	AsyncDo(func())
	IsClosed() bool
}

type BaseService struct {
	IService
	waitGroup sync.WaitGroup // wait for all goroutines
	isRunning bool
}

func (serv *BaseService) Start() error {
	serv.isRunning = true
	return nil
}

func (serv *BaseService) Close() {
	serv.isRunning = false
	serv.waitGroup.Wait()
}

func (serv *BaseService) AsyncDo(fn func()) {
	serv.waitGroup.Add(1)
	go func() {
		defer serv.waitGroup.Done()
		fn()
	}()
}

func (serv *BaseService) IsClosed() bool {
	return !serv.isRunning
}
