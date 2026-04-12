package uniqid

import (
	"sync/atomic"
)

type ID uint

type Factory struct {
	lastId atomic.Uint32
}

func New() *Factory {
	return &Factory{}
}

func (f *Factory) Next() ID {
	return ID(f.lastId.Add(1))
}
