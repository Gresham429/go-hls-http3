package control_message

import (
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
)

// Abort is the payload of the abort message
type AbortMessagePayload struct {
	ChunckStreamID uint32
}

// CreateHeader creates the header for the abort message
func (a *AbortMessagePayload) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeAbortMessage),
		Length:        15,
		StreamID:      0,
	}
}

// CreatePayload creates the payload for the abort message
func (a *AbortMessagePayload) CreatePayload() []byte {
	payload := make([]byte, 4)
	payload[0] = byte(a.ChunckStreamID >> 24)
	payload[1] = byte(a.ChunckStreamID >> 16)
	payload[2] = byte(a.ChunckStreamID >> 8)
	payload[3] = byte(a.ChunckStreamID)

	return payload
}

// CreateRTMPMessage
func (a *AbortMessagePayload) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  a.CreateHeader(),
		Payload: a.CreatePayload(),
	}
}
