package orderedmap

import "errors"

var (
	ErrKeyDoesNotExist = errors.New("key does not exist")
	ErrKeyAlreadyExist = errors.New("key already exists")
	ErrIndexOutOfRange = errors.New("index out of range")
)
