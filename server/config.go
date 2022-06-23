package server

import (
	"strconv"
	"wallet/cmd"
)

type Server struct {
	Address string
	Ssl     bool
}

func (s *Server) Configure() {
	s.Address = cmd.EnvConf.ServerHost + ":" + strconv.Itoa(cmd.EnvConf.ServerPort)
}
