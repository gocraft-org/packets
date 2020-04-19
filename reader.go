package packets

import (
	"errors"
	"io"
)

var (
	// ErrPacketLengthTooSmall means the received packet length was less than 0
	ErrPacketLengthTooSmall = errors.New("The received packet length was less than 0")
	// ErrPacketTypeTooSmall means the received packet type was less than 0
	ErrPacketTypeTooSmall = errors.New("The received packet type was less than 0")
)

// Reader is a packet reader that reads from the stream
type Reader struct {
	r io.Reader
}

// NewReader creates a new packet reader
func NewReader(r io.Reader) *Reader {
	return &Reader{r}
}
