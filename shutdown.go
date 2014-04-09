package main

import (
	"os"
	"os/signal"
)

func OnSignal(f func(), signals ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	<-c
	f()
}

func OnSignalLoop(f func(), signals ...os.Signal) {
	for {
		OnSignal(f, signals...)
	}
}
