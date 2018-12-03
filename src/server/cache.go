package server

type ICache interface {
	Insert(k, v []byte) error
	Commit() error
}
