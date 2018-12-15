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
		err = errors.New("insert param error")
		return
	}
	return cc.server.cache.Insert(data[0], data[1])
}

func (cc *clientConn) delete(data [][]byte) (err error) {
	if len(data) == 0 {
		err = errors.New("delete param error")
		return
	}
	return cc.server.cache.Delete(data[0])
}

func (cc *clientConn) update(data [][]byte) (err error) {
	if len(data) < 2 {
		err = errors.New("update param error")
		return
	}
	return cc.server.cache.Update(data[0], data[1])
}

func (cc *clientConn) get(data [][]byte) (val string, err error) {
	if len(data) == 0 {
		err = errors.New("get param error")
		return
	}
	return cc.server.cache.Get(data[0])
}
