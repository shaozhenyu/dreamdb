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

func (db *DB) Delete(k []byte) error {
	key := string(k)
	if _, ok := db.data[key]; ok {
		delete(db.data, key)
		return nil
	}
	return errors.New("key not exists")
}

func (db *DB) Update(k, v []byte) error {
	key := string(k)
	value := string(v)
	if _, ok := db.data[key]; ok {
		db.data[key] = value
		return nil
	}
	return errors.New("key not exists")
}

func (db *DB) Get(k []byte) (string, error) {
	key := string(k)
	if v, ok := db.data[key]; ok {
		return v, nil
	}
	return "", errors.New("key not exists")
}
