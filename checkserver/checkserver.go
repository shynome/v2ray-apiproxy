package checkserver

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	UUID "github.com/google/uuid"
)

// CheckServer to checkout server is ok
type CheckServer struct {
	Port            uint16
	Resp            string
	HTTPServer      *http.Server
	HTTPServerError chan error
}

// Config CheckServer
type Config struct {
	Port uint16
}

// New CheckServer
func New(config Config) (checkServer *CheckServer) {

	resp := UUID.New().String()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(resp))
	})

	checkServer = &CheckServer{
		Port: config.Port,
		Resp: resp,
		HTTPServer: &http.Server{
			Addr:    fmt.Sprintf("127.0.0.1:%v", config.Port),
			Handler: mux,
		},
		HTTPServerError: make(chan error),
	}

	return
}

func (server *CheckServer) checkIsReady(needCheckPort uint16) (chanErr chan error) {

	chanErr = make(chan error, 1)
	client := &http.Client{Timeout: time.Second}
	var checkURL string
	if needCheckPort == 0 {
		checkURL = fmt.Sprintf("http://%v", server.HTTPServer.Addr)
	} else {
		checkURL = fmt.Sprintf("http://127.0.0.1:%v", needCheckPort)
	}
	c := 0
	t := time.Duration(3 / 2)
	for ; c < 5; c++ {
		time.Sleep(t * time.Second)
		resp, err := client.Get(checkURL)
		if err != nil {
			fmt.Print(err)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			chanErr <- err
			return
		}
		if string(body) == server.Resp {
			chanErr <- nil
			return
		}
	}
	chanErr <- fmt.Errorf("check server spend too much time to start, more than %v second", time.Duration(c)*t)
	return

}

func (server CheckServer) startHTTPServer() {
	server.HTTPServerError <- server.HTTPServer.ListenAndServe()
}

// Run check server
func (server *CheckServer) Run() (err error) {

	go server.startHTTPServer()

	select {
	case err = <-server.checkIsReady(0):
	case err = <-server.HTTPServerError:
	}
	if err != nil {
		server.HTTPServer.Close()
	}

	return

}
