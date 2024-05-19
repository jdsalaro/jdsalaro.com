package main

import (
	"time"
	"jdsalaro.com/kata/codewars/opposites-attract/client"
	"jdsalaro.com/kata/codewars/opposites-attract/server"
)

/*
 *
 * code sample and playground
 * for friends guayahack.co
 *
 * by @jdsalaro
 */

//todo: use env variables

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
