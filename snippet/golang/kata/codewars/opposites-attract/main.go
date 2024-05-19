package main

import (
	"time"
	"jdsalaro.com/kata/codewars/opposites-attract/client"
	"jdsalaro.com/kata/codewars/opposites-attract/server"
)

/*
 *
 * code sample and playground
 * for friends at guayahack.co
 *
 * by @jdsalaro
 */

//todo: add a couple of tests for client.go and server.go

func main() {
	var s server.Server
	s = server.New()
	s.Run()

	time.Sleep(1 * time.Second)

	var c client.Client
	c = client.New()
	c.Run()

	select {}
}
