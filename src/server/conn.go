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
	server       *Server
	bufReadConn  net.Conn
	bufWriteConn net.Conn
}

func newClientConn(s *Server) *clientConn {
	return &clientConn{
		server: s,
	}
}

func (cc *clientConn) setConn(conn net.Conn) {
	cc.bufReadConn = conn
	cc.bufWriteConn = conn
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
		msg, err := cc.dispatch(bytes.Split(buf, []byte(" ")))
		if err != nil {
			msg = err.Error()
		}

		// send result
		_, err = cc.bufWriteConn.Write([]byte(msg))
		if err != nil {
			log.Fatalf("cc.bufWriteConn.Write error(%s)", err.Error())
		}
	}
}

func (cc *clientConn) dispatch(data [][]byte) (msg string, err error) {
	msg = "success"
	cmd := data[0]
	cmd = bytes.ToLower(cmd)
	data = data[1:]
	switch string(cmd) {
	case INSERT:
		err = cc.insert(data)
	case DELETE:
		err = cc.delete(data)
	case UPDATE:
		err = cc.update(data)
	case SELECT:
		return cc.get(data)
	default:
		err = errors.New("undefine cmd error")
	}

	if err == nil {
		err = cc.server.cache.Commit()
	}
	return
}
