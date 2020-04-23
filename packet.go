package packets

import (
	"io"

	"github.com/gocraft-org/protocol/types"
)

// Packet is a packet send back and forth that can be encoded and decoded
type Packet interface {
	Encode(w io.Writer) error
	Decode(r io.Reader) error
}

// Buffer is a buffer for storing packet information before writing to the client connection
type Buffer struct {
	d []byte
}

// Write writes data to the packet from encoding packet data
func (b *Buffer) Write(p []byte) (int, error) {
	b.d = append(b.d, p...)

	return len(p), nil
}

// Encode encodes the packet buffer into the writer for sending to the client
func (b *Buffer) Encode(t int, w io.Writer) error {
	packetType := types.VarInt(t)

	length := types.VarInt(len(b.d) + packetType.Length())

	err := length.Encode(w)

	if err != nil {
		return err
	}

	err = packetType.Encode(w)

	if err != nil {
		return err
	}

	_, err = w.Write(b.d)

	return err
}
