package control_message

import (
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
)

// AcknowledgementPayload is the payload of the acknowledgement message
type AcknowledgementPayload struct {
	SequenceNumber uint32
}

// CreateHeader creates the header for the acknowledgement message
func (a *AcknowledgementPayload) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeAcknowledgement),
		Length:        15,
		StreamID:      0,
	}
}

// CreatePayload creates the payload for the acknowledgement message
func (a *AcknowledgementPayload) CreatePayload() []byte {
	payload := make([]byte, 4)
	payload[0] = byte(a.SequenceNumber >> 24)
	payload[1] = byte(a.SequenceNumber >> 16)
	payload[2] = byte(a.SequenceNumber >> 8)
	payload[3] = byte(a.SequenceNumber)

	return payload
}

// CreateRTMPMessage
func (a *AcknowledgementPayload) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  a.CreateHeader(),
		Payload: a.CreatePayload(),
	}
}
