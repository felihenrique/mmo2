package events

import (
	"io"
	"sync"
)

type Writer struct {
	buffer     []byte
	bufferLock sync.Mutex
}

func NewWriter() *Writer {
	writer := Writer{}
	writer.buffer = make([]byte, 0)
	return &writer
}

func (w *Writer) Append(data []byte) {
	w.bufferLock.Lock()
	w.buffer = append(w.buffer, data...)
	w.bufferLock.Unlock()
}

func (w *Writer) Send(writer io.Writer) (int64, error) {
	w.bufferLock.Lock()
	oldBuffer := w.buffer
	w.buffer = make([]byte, 0)
	w.bufferLock.Unlock()
	n, err := writer.Write(oldBuffer)
	if n < len(oldBuffer) {
		w.bufferLock.Lock()
		w.buffer = append(oldBuffer[n:], w.buffer...)
		w.bufferLock.Unlock()
	}
	return int64(n), err
}
