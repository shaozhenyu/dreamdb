package server

import (
	"fmt"
	"net"
)

type Server struct {
	listener net.Listener
	cache    ICache
}

func NewServer(cache ICache) (*Server, error) {
	s := &Server{cache: cache}

	var err error
	ip := "127.0.0.1"
	port := 8081
	addr := fmt.Sprintf("%s:%d", ip, port)
	if s.listener, err = net.Listen("tcp", addr); err == nil {
		fmt.Printf("Server is Running at [%s]\n", addr)
	}
	return s, err
}

func (s *Server) Run() error {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		go s.onConn(conn)
	}
	err := s.listener.Close()
	if err != nil {
		return err
	}
	s.listener = nil
	return nil
}

func (s *Server) newConn(conn net.Conn) *clientConn {
	cc := newClientConn(s)
	cc.setConn(conn)
	return cc
}

func (s *Server) onConn(c net.Conn) {
	defer c.Close()
	conn := s.newConn(c)
	conn.Run()
}
