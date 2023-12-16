package gsp

import (
	"context"
	"errors"
	"io"
	"net"
	"os"
	"syscall"
)

var ErrDisconnected = errors.New("client disconnected")
var ErrTimeout = errors.New("IO timeout")

func handleError(err error) error {
	e, ok := err.(net.Error)
	isTimeout :=
		ok &&
			e.Timeout() ||
			errors.Is(err, os.ErrDeadlineExceeded) ||
			errors.Is(err, context.DeadlineExceeded)

	if isTimeout {
		return ErrTimeout
	}

	isDisconnected :=
		errors.Is(err, syscall.ECONNRESET) ||
			errors.Is(err, syscall.ECONNABORTED) ||
			errors.Is(err, syscall.EPIPE) ||
			errors.Is(err, io.ErrClosedPipe) ||
			errors.Is(err, io.EOF)

	if isDisconnected {
		return ErrDisconnected
	}

	return err
}
