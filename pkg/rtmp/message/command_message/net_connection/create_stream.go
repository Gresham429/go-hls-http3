package net_connection

import (
	"bytes"

	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message/command_message"
	amf "github.com/elobuff/goamf"
)

// CreateStreamPayloadAMF0 is the payload of the createStream command message
type CreateStreamResponsePayloadAMF0 struct {
	command_message.CommandMessagePayload
	StreamID uint32
}

// CreateStreamCommandName is the command name for the createStream command message
const (
	CreateStreamCommandName = "createStream"
	CreateStreamResult      = "_result"
	CreateStreamError       = "_error"
)

// CreateHeader creates a header for the createStream command message
func (createStream *CreateStreamResponsePayloadAMF0) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeCommandMessageAMF0),
		Length:        uint32(11 + len(createStream.CreatePayload())),
		StreamID:      0,
	}
}

// CreatePayload creates a payload for the createStream command message
func (createStream *CreateStreamResponsePayloadAMF0) CreatePayload() []byte {
	var payload []byte
	buf := new(bytes.Buffer)

	encoder := new(amf.Encoder)

	_, err := encoder.EncodeAmf0(buf, createStream.CommandName)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0(buf, createStream.TransactionID)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0Object(buf, createStream.CommandObject, true)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0(buf, createStream.StreamID)
	if err != nil {
		return nil
	}

	payload = append(payload, buf.Bytes()...)

	return payload
}

// CreateRTMPMessage creates a RTMP message for the createStream command message
func (createStream *CreateStreamResponsePayloadAMF0) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  createStream.CreateHeader(),
		Payload: createStream.CreatePayload(),
	}
}

// CreateStreamRequestPayloadAMF3 is the payload of the createStream command message
type CreateStreamRequestPayloadAMF3 struct {
	command_message.CommandMessagePayload
	StreamID uint32
}

// CreateHeader creates a header for the createStream command message
func (createStream *CreateStreamRequestPayloadAMF3) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeCommandMessageAMF3),
		Length:        uint32(1 + len(createStream.CreatePayload())),
		StreamID:      0,
	}
}

// CreatePayload creates a payload for the createStream command message
func (createStream *CreateStreamRequestPayloadAMF3) CreatePayload() []byte {
	var payload []byte
	buf := new(bytes.Buffer)

	encoder := new(amf.Encoder)

	_, err := encoder.EncodeAmf3(buf, createStream.CommandName)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf3(buf, createStream.TransactionID)
	if err != nil {
		return nil
	}

	to := *amf.NewTypedObject()
	to.Type = "org.amf.ASClass"
	to.Object = createStream.CommandObject

	_, err = encoder.EncodeAmf3Object(buf, to, true)
	if err != nil {
		return nil
	}


	payload = append(payload, buf.Bytes()...)

	return payload
}
