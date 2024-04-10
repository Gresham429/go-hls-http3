// message.go
package message

// IMessage interface defines the basic structure of a message
type MessageInterface interface {
	GetHeader() *RTMPHeader
	GetPayload() []byte
	GetPayloadLength() uint32
	CreateHeader() *RTMPHeader
	CreatePayload() []byte
	CreateRTMPMessage() *RTMPMessage
}

// MessageType is the type of the message
type MessageType int

// RTMP message types
const (
	MessageTypeSetChunkSize MessageType = iota + 1
	MessageTypeAbortMessage
	MessageTypeAcknowledgement
	MessageTypeUserControlMessage
	MessageTypeWindowAcknowledgementSize
	MessageTypeSetPeerBandwidth

	MessageTypeDataMessageAMF3         = 15
	MessageTypeSharedObjectMessageAMF3 = 16
	MessageTypeCommandMessageAMF3      = 17
	MessageTypeDataMessageAMF0         = 18
	MessageTypeSharedObjectMessageAMF0 = 19
	MessageTypeCommandMessageAMF0      = 20

	MessageTypeAudioMessage = 8
	MessageTypeVideoMessage = 9

	MessageTypeAggregateMessage = 22
)

// RTMPMessage is the structure of the RTMP message
type RTMPMessage struct {
	Header  *RTMPHeader
	Payload []byte // Payload is the data of the message
}

// RTMPHeader is the structure of the RTMP header
type RTMPHeader struct {
	Timestamp     uint32
	MessageTypeID uint8
	Length        [3]byte
	StreamID      [3]byte
}

// GetHeader returns the header of the RTMP message
func (m *RTMPMessage) GetHeader() *RTMPHeader {
	return m.Header
}

// GetPayload returns the payload of the RTMP message
func (m *RTMPMessage) GetPayload() []byte {
	return m.Payload
}
