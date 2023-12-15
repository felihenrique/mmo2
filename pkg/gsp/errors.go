package gsp

import (
	"errors"
	"io"
	"syscall"
)

var ErrDisconnected = errors.New("client disconnected")

func handleError(err error) error {
	switch err {
	case syscall.ECONNRESET, syscall.ECONNABORTED, syscall.EPIPE, io.ErrClosedPipe, io.EOF:
		return ErrDisconnected
	default:
		return err
	}
}
