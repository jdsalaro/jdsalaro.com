
package client

import (
	"io"
	"fmt"
	"log"
	"net/http"
)

type ClientConfig struct {
	ConnectRetries int
	Url string
	Port int
}

type Client struct {
	Config ClientConfig
}

type ClientOption func(c *ClientConfig)

func WithPort(port int) ClientOption {
	return func(c *ClientConfig) {
		c.Port = port
	}
}
func WithConnectRetries(retries int) ClientOption {
	return func(c *ClientConfig) {
		c.ConnectRetries = retries
	}
}

func WithUrl(url string) ClientOption {
	return func(c *ClientConfig) {
		c.Url = url
	}
}

func New(opts ... ClientOption) Client {
	config := &ClientConfig{
		ConnectRetries: 10,
		Port:8080,
		Url: fmt.Sprintf(
				"http://localhost:%d?%s",
				8080,
				"sarah=4&timmy=5",
			),
	}
	for _, opt := range opts {
		opt(config)
	}
	return Client{
		Config: *config,
	}
}

func (c Client) Run() error {
	log.Println("RUNNING CLIENT")

	// timmyn:=rand.Intn(10)
	// sarahn:=rand.Intn(10)

	// url := fmt.Sprintf("http://localhost:%d?sarah=%d&timmy=%d", port, sarahn, timmyn)
	
	err := *new(error)
	err = error(nil)	//equivalent to the above

	var client http.Client
	resp, err := client.Get(c.Config.Url)
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
		log.Printf("VERDICT: %s", bodyString)
	}

	return err
}


