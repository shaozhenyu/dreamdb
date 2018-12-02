package db

import (
	"errors"
)

func (db *DB) Insert(k, v []byte) error {
	key := string(k)
	value := string(v)
	if _, ok := db.data[key]; ok {
		return errors.New("key has exists")
	}
	db.data[string(k)] = value
	return nil
}
