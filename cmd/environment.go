package cmd

import (
	"os"
	"strconv"
)

type Env struct {
	ServerHost string
	ServerPort int
}

var EnvConf *Env = &Env{}

func (e *Env) ParseEnvVars() {
	e.ServerHost = os.Getenv("SERVER_HOST")
	e.ServerPort, _ = strconv.Atoi(os.Getenv("SERVER_PORT"))
}
