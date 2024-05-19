
package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
)

type ServerConfig struct {
	Port int
}

type Server struct {
	Config ServerConfig
}

type ServerOption func(c *ServerConfig)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	inLove := false

	if len(params) == 2 {
		timmy := params.Get("timmy")
		tn, _ := strconv.Atoi(timmy)
		log.Printf("TIMMY PETALS: %d", tn)
		sarah := params.Get("sarah")
		sn, _ := strconv.Atoi(sarah)
		log.Printf("SARAH PETALS: %d", sn)
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

func New(opts ... ServerOption) Server {
	config := &ServerConfig{
		Port:8080,
	}
	for _, opt := range opts {
		opt(config)
	}
	return Server{
		Config: *config,
	}
}

func (s Server) Run() error {
	log.Println("RUNNING SERVER")
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Config.Port))
	if err != nil {
		panic(err)
	}
	h := &http.Server{
		Handler: http.HandlerFunc(DefaultHandler),
	}

	go h.Serve(l)
	return err
}