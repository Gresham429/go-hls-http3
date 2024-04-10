package control_message

import (
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
	"github.com/Gresham429/go-hls-http3/utils"
)

// SetPeerBandwidthPayload is the payload of the SetPeerBandwidth message
type SetPeerBandwidthPayload struct {
	WindowSize uint32
	LimitType  uint8
}

// CreateHeader creates the header for the SetPeerBandwidth message
func (s *SetPeerBandwidthPayload) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeSetPeerBandwidth),
		Length:        utils.ConvertToThreeBytes(11),
		StreamID:      utils.ConvertToThreeBytes(0),
	}
}

// CreatePayload creates the payload for the SetPeerBandwidth message
func (s *SetPeerBandwidthPayload) CreatePayload() []byte {
	payload := make([]byte, 5)
	payload[0] = byte(s.WindowSize >> 24)
	payload[1] = byte(s.WindowSize >> 16)
	payload[2] = byte(s.WindowSize >> 8)
	payload[3] = byte(s.WindowSize)
	payload[4] = s.LimitType

	return payload
}

// CreateRTMPMessage
func (s *SetPeerBandwidthPayload) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  s.CreateHeader(),
		Payload: s.CreatePayload(),
	}
}