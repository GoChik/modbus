package modbus

import (
	"io"
)

type serial struct {
	transport io.ReadWriteCloser
}

func (s *serial) Read(p []byte) (n int, err error) {
	return s.transport.Read(p)
}

func (s *serial) Write(p []byte) (n int, err error) {
	return s.transport.Write(p)
}

func (s *serial) Close() error {
	return s.transport.Close()
}
