package api

import ("github.com/deepakdinesh1123/mezcala/pkgs/agent"
"github.com/nats-io/nats.go"
)

type Server struct {
	agent *agent.Agent
	nats.go
}

func NewServer() {

}
