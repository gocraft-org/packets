package packets

import (
	"github.com/gocraft-org/protocol/types"
)

// HandshakePacket is a Handshake (0x00) sent to the server from the client
type HandshakePacket struct {
	ProtocolVersion types.VarInt
	ServerAddress   types.String
	ServerOirt      types.UnsignedShort
	NextState       types.VarInt
}
