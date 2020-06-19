package main

import (
	"context"
	"fmt"
	"net"
)

type Server struct {
	Info *ServerInfo
	Sock net.Listener
	Context context.Context
}

type ServerInfo struct {
	Addr string
	Port string
	Proto string
}

func ServerInit(info *ServerInfo, ctx context.Context) (*Server, error) {
	host := fmt.Sprintf("%s:%s", info.Addr, info.Port)
	sock, err := net.Listen(info.Proto, host)
	if err != nil {
		return nil, err
	}
	return &Server{
		Info: info,
		Sock: sock,
		Context: ctx,
	}, nil
}