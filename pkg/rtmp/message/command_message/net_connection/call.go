package net_connection

import (
	"bytes"

	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message/command_message"
	amf "github.com/elobuff/goamf"
)

// CallResponsePayloadAMF0 is the payload of the call command message
type CallResponsePayloadAMF0 struct {
	command_message.CommandMessagePayload
	Response amf.Object
}

const (
	CallCommandName = "call"
)

// CreateHeader creates a header for the call command message
func (call *CallResponsePayloadAMF0) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeCommandMessageAMF0),
		Length:        uint32(11 + len(call.CreatePayload())),
		StreamID:      0,
	}
}

// CreatePayload creates a payload for the call command message
func (call *CallResponsePayloadAMF0) CreatePayload() []byte {
	var payload []byte
	buf := new(bytes.Buffer)
	encoder := new(amf.Encoder)

	_, err := encoder.EncodeAmf0(buf, call.CommandName)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0(buf, call.TransactionID)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0Object(buf, call.CommandObject, true)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0Object(buf, call.Response, true)
	if err != nil {
		return nil
	}

	payload = append(payload, buf.Bytes()...)

	return payload
}

// CreateRTMPMessage creates a RTMP message for the call command message
func (call *CallResponsePayloadAMF0) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  call.CreateHeader(),
		Payload: call.CreatePayload(),
	}
}

// CallResponsePayloadAMF3 is the payload of the call command message
type CallResponsePayloadAMF3 struct {
	command_message.CommandMessagePayload
	Response amf.Object
}

// CreateHeader creates a header for the call command message
func (call *CallResponsePayloadAMF3) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeCommandMessageAMF3),
		Length:        uint32(7 + len(call.CreatePayload())),
		StreamID:      0,
	}
}

// CreatePayload creates a payload for the call command message
func (call *CallResponsePayloadAMF3) CreatePayload() []byte {
	var payload []byte
	buf := new(bytes.Buffer)
	encoder := new(amf.Encoder)

	_, err := encoder.EncodeAmf3(buf, call.CommandName)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf3(buf, call.TransactionID)
	if err != nil {
		return nil
	}

	to := *amf.NewTypedObject()
	to.Type = "org.amf.ASClass"
	to.Object = call.CommandObject

	_, err = encoder.EncodeAmf3Object(buf, to, true)
	if err != nil {
		return nil
	}

	to.Object = call.Response
	_, err = encoder.EncodeAmf3Object(buf, to, true)
	if err != nil {
		return nil
	}

	payload = append(payload, buf.Bytes()...)

	return payload
}
