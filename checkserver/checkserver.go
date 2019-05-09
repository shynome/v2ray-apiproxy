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

	mux := http.NewServeMux()

	checkServer = &CheckServer{
		Port: config.Port,
		Resp: UUID.New().String(),
		HTTPServer: &http.Server{
			Addr:    fmt.Sprintf("127.0.0.1:%v", config.Port),
			Handler: mux,
		},
		HTTPServerError: make(chan error),
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(checkServer.Resp))
	})

	return
}

// CheckPortIsReady check port
func CheckPortIsReady(needCheckPort uint16, Resp string, maxTryTimes uint) (chanErr chan error) {

	chanErr = make(chan error, 1)
	client := &http.Client{Timeout: time.Second}
	var checkURL = fmt.Sprintf("http://127.0.0.1:%v", needCheckPort)
	c := uint(0)
	t := time.Duration(3 / 2)
	for ; c < maxTryTimes; c++ {
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
		if string(body) == Resp {
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
	case err = <-CheckPortIsReady(server.Port, server.Resp, 5):
	case err = <-server.HTTPServerError:
	}
	if err != nil {
		server.HTTPServer.Close()
	}

	return

}

// Close check server
func (server *CheckServer) Close() (err error) {
	return server.HTTPServer.Close()
}
