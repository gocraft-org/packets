package packet

import (
	"io"

	"github.com/gocraft-org/protocol/types"
)

// StatusResponsePacket is a Status Response packet (0x00) sent to the client from the server
type StatusResponsePacket struct {
	JSONData types.String
}

// Encode encodes the packet into the writer
func (p *StatusResponsePacket) Encode(w io.Writer) error {
	err := p.JSONData.Encode(w)

	if err != nil {
		return err
	}

	return nil
}

// Decode decodes the reader into the packet
func (p *StatusResponsePacket) Decode(r io.Reader) error {
	err := p.JSONData.Decode(r)

	if err != nil {
		return err
	}

	return nil
}

// StatusPongPacket is a Status Pong packet (0x01) sent to the client from the server
type StatusPongPacket struct {
	Payload types.Long
}

// Encode encodes the packet into the writer
func (p *StatusPongPacket) Encode(w io.Writer) error {
	err := p.Payload.Encode(w)

	if err != nil {
		return err
	}

	return nil
}

// Decode decodes the reader into the packet
func (p *StatusPongPacket) Decode(r io.Reader) error {
	err := p.Payload.Decode(r)

	if err != nil {
		return err
	}

	return nil
}

// StatusRequestPacket is a Status Request packet (0x00) sent to the server from the client
type StatusRequestPacket struct{}

// StatusPingPacket is a Status Ping packet (0x01) sent to the server from the client
type StatusPingPacket struct {
	Payload types.Long
}

// Encode encodes the packet into the writer
func (p *StatusPingPacket) Encode(w io.Writer) error {
	err := p.Payload.Encode(w)

	if err != nil {
		return err
	}

	return nil
}

// Decode decodes the reader into the packet
func (p *StatusPingPacket) Decode(r io.Reader) error {
	err := p.Payload.Decode(r)

	if err != nil {
		return err
	}

	return nil
}
