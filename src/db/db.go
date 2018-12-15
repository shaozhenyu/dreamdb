package db

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type DB struct {
	path string
	file *os.File
	data map[string]string
}

func Open(path string, mode os.FileMode) (db *DB, err error) {
	defer func() {
		if db != nil && err != nil {
			db.close()
		}
	}()

	db = &DB{
		path: path,
		data: make(map[string]string, 4096),
	}
	db.file, err = os.OpenFile(path, os.O_CREATE|os.O_RDWR, mode)
	if err != nil {
		return nil, err
	}
	if info, err := db.file.Stat(); err != nil {
		return nil, err
	} else if info.Size() != 0 {
		if err = db.init(); err != nil {
			return nil, err
		}
	}
	return
}

func (db *DB) String() string {
	return fmt.Sprintf("DB<%q>", db.path)
}

func (db *DB) Commit() error {
	// truncate
	if err := db.file.Truncate(0); err != nil {
		return err
	}

	var buf bytes.Buffer
	for k, v := range db.data {
		buf.WriteString(k)
		buf.WriteRune(' ')
		buf.WriteString(v)
		buf.WriteRune('\n')
	}
	_, err := db.file.WriteAt(buf.Bytes(), 0)
	return err
}

func (db *DB) init() error {
	if db.data == nil {
		db.data = make(map[string]string, 4096)
	}
	reader := bufio.NewReader(db.file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		kv := bytes.Split(line, []byte(" "))
		db.data[string(kv[0])] = string(kv[1])
	}
	return nil
}

func (db *DB) close() {
	if db != nil {
		db.file.Close()
	}
}
