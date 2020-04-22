package packets

import (
	"io"

	"github.com/gocraft-org/protocol/types"
)

// HandshakePacket is a Handshake (0x00) sent to the server from the client
type HandshakePacket struct {
	ProtocolVersion types.VarInt
	ServerAddress   types.String
	ServerPort      types.UnsignedShort
	NextState       types.VarInt
}

// Encode encodes the packet into the writer
func (p *HandshakePacket) Encode(w io.Writer) error {
	err := p.ProtocolVersion.Encode(w)

	if err != nil {
		return err
	}

	err = p.ServerAddress.Encode(w)

	if err != nil {
		return err
	}

	err = p.ServerPort.Encode(w)

	if err != nil {
		return err
	}

	err = p.NextState.Encode(w)

	if err != nil {
		return err
	}

	return nil
}

// Decode decodes the reader into the packet
func (p *HandshakePacket) Decode(r io.Reader) error {
	err := p.ProtocolVersion.Decode(r)

	if err != nil {
		return err
	}

	err = p.ServerAddress.Decode(r)

	if err != nil {
		return err
	}

	err = p.ServerPort.Decode(r)

	if err != nil {
		return err
	}

	err = p.NextState.Decode(r)

	if err != nil {
		return err
	}

	return nil
}
