
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

const (
	PARAMS_EXPECTED = 2
	PORT_LISTEN = 2
)

func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	inLove := false

	if len(params) == PARAMS_EXPECTED {

		pn := []int{}

		for k, v := range params {
			n, _ := strconv.Atoi(v[0])
			pn = append(pn,n )
			log.Printf("PETALS: %s WITH %d", k, n)
		}

		if (pn[0]+pn[1])%2 != 0 {
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
		Port:PORT_LISTEN,
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