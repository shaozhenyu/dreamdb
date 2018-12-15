package server

type ICache interface {
	Insert(k, v []byte) error
	Delete(k []byte) error
	Update(k, v []byte) error
	Get(k []byte) (string, error)
	Commit() error
}
