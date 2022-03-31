package main

import (
	"fiveGCHub/common"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const HubVersion = "0.0.1"

type MethodSet struct {
	ms      map[string]common.Method
	version string
}

func NewMethodSet() *MethodSet {
	ms := common.GetMethodSet()
	return &MethodSet{ms, HubVersion}
}

func (m *MethodSet) MethodRun(t *time.Ticker, method common.Method, stopCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-t.C:
			method.Run()
		case <-stopCh:
			return
		}
	}
}

func (m *MethodSet) Run(stopCh <-chan struct{}, wg *sync.WaitGroup) {
	for _, ms := range m.ms {
		t := time.NewTicker(time.Duration(ms.GetInterval()) * time.Second)
		wg.Add(1)
		go m.MethodRun(t, ms, stopCh, wg)
	}
}

func main() {
	m := NewMethodSet()
	stopCh := make(chan struct{})
	wg := &sync.WaitGroup{}

	m.Run(stopCh, wg)

	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	close(stopCh)
	wg.Wait()
}
