package events

import (
	"io"
	"mmo2/pkg/serialization"
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

func (w *Writer) Append(bytes []byte) {
	eventSize := serialization.AppendInt16([]byte{}, int16(len(bytes)+2))
	bytes = append(eventSize, bytes...)
	w.bufferLock.Lock()
	w.buffer = append(w.buffer, bytes...)
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
