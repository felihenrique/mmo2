package gsp

import (
	"errors"
	"io"
	"os"
	"syscall"
)

var ErrDisconnected = errors.New("client disconnected")

func handleError(err error) error {
	switch err {
	case os.ErrDeadlineExceeded:
	case syscall.ECONNRESET:
	case syscall.ECONNABORTED:
	case syscall.EPIPE:
	case io.ErrClosedPipe:
	case io.EOF:
		return ErrDisconnected
	}
	return err
}
