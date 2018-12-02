package server

import (
	"errors"
)

const (
	INSERT = "insert"
	UPDATE = "update"
	SELECT = "select"
	DELETE = "delete"
)

func (cc *clientConn) insert(data [][]byte) (err error) {
	if len(data) < 2 {
		err = errors.New("input error")
		return
	}
	return cc.server.cache.Insert(data[0], data[1])
}
