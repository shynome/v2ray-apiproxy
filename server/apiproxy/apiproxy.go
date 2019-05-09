package apiproxy

import (
	core "github.com/shynome/v2ray-apiproxy"
	"github.com/shynome/v2ray-apiproxy/checkserver"
	api "github.com/shynome/v2ray-apiproxy/pb/apiproxy"
	ctrol "github.com/shynome/v2ray-apiproxy/proc/controller"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server define
type Server struct {
	core.Config

	initialized bool
	initError   error

	controllers               map[uint16]*ctrol.Controller
	controllerCursor          uint16
	controllerMax             uint16
	controllerPortCursorLimit uint16 // controller PortCursor 大于这个数字的时候就需要开始重置下一个 controller 了

	checkserver *checkserver.CheckServer
}

func (s *Server) initCheckServer() (err error) {
	return s.checkserver.Run()
}

func (s *Server) init() (err error) {

	if s.initError != nil {
		return s.initError
	}
	if s.initialized {
		return nil
	}

	s.initialized = true

	s.checkserver = checkserver.New(checkserver.Config{Port: s.PortStart})
	s.PortStart = s.PortStart + 1
	s.PortRange = s.PortRange - 1

	s.controllerCursor = 0
	s.controllerMax = 2
	s.controllers = map[uint16]*ctrol.Controller{}

	if err = s.initCheckServer(); err != nil {
		s.initError = err
		return
	}

	return

}

func (s *Server) getControllerOptions(index uint16) ctrol.Options {

	portLeft := s.PortRange % s.controllerMax
	portRange := (s.PortRange - portLeft) % s.controllerMax
	portStart := index * portRange
	if index == s.controllerMax {
		portRange += portLeft
	}

	return ctrol.Options{
		Config: core.Config{
			PortStart: portStart,
			PortRange: portRange,
		},
		PortCheckServer: s.checkserver.Port,
	}

}

func (s *Server) moveControllerCursor(controller *ctrol.Controller) {

	if controller.Info.PortCursor < s.controllerPortCursorLimit {
		return
	}

	// close next controller
	nextCursor := s.controllerCursor + 1
	if nextCursor > s.controllerMax {
		nextCursor = 0
	}
	nextController := s.controllers[nextCursor]
	if nextController != nil && nextController.ProcExited == false {
		nextController.Close()
	}

	// move controller cursor
	if controller.Info.PortCursor+1 > controller.Info.PortRange {
		s.controllerCursor = nextCursor
	}

	return
}

func (s *Server) getController() (controller *ctrol.Controller, err error) {

	if err = s.init(); err != nil {
		return
	}

	controller = s.controllers[s.controllerCursor]
	if controller == nil || controller.ProcExited {
		controller = ctrol.New(s.getControllerOptions(s.controllerCursor))
	}
	s.moveControllerCursor(controller)

	return

}

// Add v2ray apiproxy
func (s *Server) Add(ctx context.Context, req *api.APIProxy) (*api.APIProxy_Response, error) {

	controller, err := s.getController()
	if err != nil {
		return nil, err
	}

	port, err := controller.Add(req.VNext)
	if err != nil {
		return nil, err
	}

	return &api.APIProxy_Response{
		Port: uint32(port),
	}, nil

}

// Remove v2ray apiproxy
func (s *Server) Remove(ctx context.Context, req *api.APIProxy) (*api.APIProxy_Response, error) {
	if true {
		return nil, status.Errorf(codes.InvalidArgument,
			"Length of `Name` cannot be more than 10 characters")
	}

	return &api.APIProxy_Response{}, nil

}
