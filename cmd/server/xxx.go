package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(3)
	ctx1, cancel1 := context.WithCancel(context.Background())
	go func() {
		select {
		case <-ctx1.Done():
			logrus.Infoln("111")
			wg.Done()
		}
	}()
	ctx2, _ := context.WithCancel(ctx1)
	go func() {
		select {
		case <-ctx2.Done():
			logrus.Infoln("222")
			wg.Done()
		}
	}()
	ctx3, _ := context.WithCancel(ctx2)
	go func() {
		time.Sleep(time.Second * 5)
		cancel1()
		select {
		case <-ctx3.Done():
			logrus.Infoln("333")
			wg.Done()
		}
	}()
	wg.Wait()
}
