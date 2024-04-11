package control_message

import (
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
)

// SetChunckSizePayload is the payload of the SetChunkSize message
type SetChunckSizePayload struct {
	ChunkSize uint32
}

// CreateHeader creates the header for the SetChunkSize message
func (s *SetChunckSizePayload) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeSetChunkSize),
		Length:        15,
		StreamID:      0,
	}
}

// CreatePayload creates the payload for the SetChunkSize message
func (s *SetChunckSizePayload) CreatePayload() []byte {
	// Create a byte slice of length 4, maxsize of chunk size is 0x7FFFFFFF
	s.ChunkSize = s.ChunkSize & 0x7FFFFFFF

	payload := make([]byte, 4)
	payload[0] = byte(s.ChunkSize >> 24)
	payload[1] = byte(s.ChunkSize >> 16)
	payload[2] = byte(s.ChunkSize >> 8)
	payload[3] = byte(s.ChunkSize)

	return payload
}

// CreateRTMPMessage
func (s *SetChunckSizePayload) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  s.CreateHeader(),
		Payload: s.CreatePayload(),
	}
}
