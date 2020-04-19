package packets

import "github.com/gocraft-org/protocol/types"

// StatusResponsePacket is a Status Response packet (0x00) sent to the client from the server
type StatusResponsePacket struct {
	JSONData types.String
}

// StatusPongPacket is a Status Pong packet (0x01) sent to the client from the server
type StatusPongPacket struct {
	JSONData types.String
}

// StatusRequestPacket is a Status Request packet (0x00) sent to the server from the client
type StatusRequestPacket struct{}

// StatusPingPacket is a Status Ping packet (0x01) sent to the server from the client
type StatusPingPacket struct {
	Payload types.Long
}
