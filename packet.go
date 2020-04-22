package packets

import "io"

// Packet is a packet send back and forth that can be encoded and decoded
type Packet interface {
	Encode(w io.Writer) error
	Decode(r io.Reader) error
}
