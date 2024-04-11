package control_message

import (
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
)

// User Control Message types
const (
	UserControlMessageStreamBegin = iota
	UserControlMessageStreamEOF
	UserControlMessageStreamDry
	UserControlMessageSetBufferLength
	UserControlMessageStreamIsRecorded

	UserControlMessagePingRequest  = 6
	UserControlMessagePingResponse = 7
)

type UserControlPayload struct {
	EventType uint16
	EventData []byte
}

// GetPayloadLength
func (u *UserControlPayload) GetPayloadLength() uint32 {
	return uint32(len(u.EventData) + 2)
}

// CreateHeader creates the header for the user control message
func (u *UserControlPayload) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeUserControlMessage),
		Length:        11 + u.GetPayloadLength(),
		StreamID:      0,
	}
}

// CreatePayload creates the payload for the user control message
func (u *UserControlPayload) CreatePayload() []byte {
	payload := make([]byte, 2+len(u.EventData))
	payload[0] = byte(u.EventType >> 8)
	payload[1] = byte(u.EventType)
	copy(payload[2:], u.EventData)
	return payload
}

// CreateRTMPMessage
func (u *UserControlPayload) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  u.CreateHeader(),
		Payload: u.CreatePayload(),
	}
}
