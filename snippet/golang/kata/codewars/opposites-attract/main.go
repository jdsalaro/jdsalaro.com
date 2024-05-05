package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"time"
)

/*
 *
 * code sample and playground
 * for friends guayahack.co
 *
 * by @jdsalaro
 */

//todo: dockerize
//todo: use env variables
//todo: add quick remarks about building and running
//todo: modularize?
//todo: improve logging?
//todo: other fun stuff for practice

func handler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	inLove := false
	if len(params) == 2 {
		timmy := params.Get("timmy")
		tn, _ := strconv.Atoi(timmy)
		sarah := params.Get("sarah")
		sn, _ := strconv.Atoi(sarah)
		if (tn+sn)%2 != 0 {
			inLove = true
		}
	}
	body, _ := json.Marshal(
		map[string]bool{
			"inLove": inLove,
		},
	)

	w.Write([]byte(body))
}

func server(port int) error {
	log.Println("Starting server ...")
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	s := &http.Server{
		Handler: http.HandlerFunc(handler),
	}

	go s.Serve(l)
	return err
}

func client(port int) error {
	log.Println("Sending request ...")
	//todo: replace hardcoded numbers with rand
	url := fmt.Sprintf("http://localhost:%d?timmy=1&sarah=2", port)
	var client http.Client
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}

	return nil
}

func main() {
	server(8080)
	time.Sleep(1 * time.Second)
	client(8080)
	// select {}
}
