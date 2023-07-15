/*
Copyright © 2023 Pierre PELOILLE <pierre@peloille.com>
*/
package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pale-whale/share.me/cmd"
	"github.com/pale-whale/share.me/internal/archive"
)

func catchSignal() {
	terminateSignals := make(chan os.Signal, 1)
	reloadSignals := make(chan os.Signal, 1)

	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGTERM)

	signal.Notify(reloadSignals, syscall.SIGUSR1, syscall.SIGHUP)

	for {
		select {
		case s := <-terminateSignals:
			log.Println("Caught", s, "termitating gracefully")
			archive.RemoveAll()
			os.Exit(0)
			break
		case <-reloadSignals:
			log.Println("Will reload config some day soon™")
		}
	}
}

func main() {
	go catchSignal()
	cmd.Execute()
}
