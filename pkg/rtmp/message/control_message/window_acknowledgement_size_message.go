package control_message

import (
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
)

// WindowAcknowledgementSizePayload is the payload of the window acknowledgement size message
type WindowAcknowledgementSizePayload struct {
	WindowSize uint32
}

// CreateHeader creates the header for the window acknowledgement size message
func (w *WindowAcknowledgementSizePayload) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeWindowAcknowledgementSize),
		Length:        15,
		StreamID:      0,
	}
}

// CreatePayload creates the payload for the window acknowledgement size message
func (w *WindowAcknowledgementSizePayload) CreatePayload() []byte {
	payload := make([]byte, 4)
	payload[0] = byte(w.WindowSize >> 24)
	payload[1] = byte(w.WindowSize >> 16)
	payload[2] = byte(w.WindowSize >> 8)
	payload[3] = byte(w.WindowSize)

	return payload
}

// CreateRTMPMessage
func (w *WindowAcknowledgementSizePayload) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  w.CreateHeader(),
		Payload: w.CreatePayload(),
	}
}
