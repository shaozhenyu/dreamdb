package server

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type clientConn struct {
	server      *Server
	bufReadConn net.Conn
}

func newClientConn(s *Server) *clientConn {
	return &clientConn{
		server: s,
	}
}

func (cc *clientConn) setConn(conn net.Conn) {
	cc.bufReadConn = conn
}

func (cc *clientConn) Run() {
	for {
		buf := make([]byte, 1024)
		size, err := cc.bufReadConn.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Conn End")
				break
			} else {
				log.Fatalf("conn.Read error(%s)", err.Error())
			}
		}
		buf = buf[:size]
		err = cc.dispatch(bytes.Split(buf, []byte(" ")))
		if err != nil {
			log.Fatalf("cc.dispatch error(%s)", err.Error())
		}
	}
}

func (cc *clientConn) dispatch(data [][]byte) (err error) {
	fmt.Println("dispatch:", string(data[0]))
	cmd := data[0]
	cmd = bytes.ToLower(cmd)
	data = data[1:]
	switch string(cmd) {
	case INSERT:
		err = cc.insert(data)
	default:
		err = errors.New("undefine cmd error")
	}

	if err == nil {
		err = cc.server.cache.Commit()
	}
	return
}
