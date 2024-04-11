package net_connection

import (
	"bytes"

	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message"
	"github.com/Gresham429/go-hls-http3/pkg/rtmp/message/command_message"
	amf "github.com/elobuff/goamf"
)

// ConnectPayloadAMF0 is the payload of the connect command message
type ConnectPayloadAMF0 struct {
	command_message.CommandMessagePayload
	OptionalUserArguments amf.Object
}

const (
	ConnectCommandName = "connect"
)

// CreateHeader creates a header for the connect command message
func (connect *ConnectPayloadAMF0) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeCommandMessageAMF0),
		Length:        uint32(11 + len(connect.CreatePayload())),
		StreamID:      0,
	}
}

// CreatePayload creates a payload for the connect command message
func (connect *ConnectPayloadAMF0) CreatePayload() []byte {
	var payload []byte
	buf := new(bytes.Buffer)
	encoder := new(amf.Encoder)

	_, err := encoder.EncodeAmf0(buf, connect.CommandName)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0(buf, connect.TransactionID)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0Object(buf, connect.CommandObject, true)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf0Object(buf, connect.OptionalUserArguments, true)
	if err != nil {
		return nil
	}

	payload = append(payload, buf.Bytes()...)

	return payload
}

// CreateRTMPMessage creates a RTMP message for the connect command message
func (connect *ConnectPayloadAMF0) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  connect.CreateHeader(),
		Payload: connect.CreatePayload(),
	}
}

// ConnectPayloadAMF3 is the payload of the connect command message
type ConnectPayloadAMF3 struct {
	command_message.CommandMessagePayload
	OptionalUserArguments map[string]interface{}
}

// CreateHeader creates a header for the connect command message
func (connect *ConnectPayloadAMF3) CreateHeader() *message.RTMPHeader {
	return &message.RTMPHeader{
		MessageTypeID: uint8(message.MessageTypeCommandMessageAMF3),
		Length:        uint32(11 + len(connect.CreatePayload())),
		StreamID:      0,
	}
}

// CreatePayload creates a payload for the connect command message
func (connect *ConnectPayloadAMF3) CreatePayload() []byte {
	var payload []byte
	buf := new(bytes.Buffer)
	encoder := new(amf.Encoder)

	_, err := encoder.EncodeAmf3(buf, connect.CommandName)
	if err != nil {
		return nil
	}

	_, err = encoder.EncodeAmf3(buf, connect.TransactionID)
	if err != nil {
		return nil
	}

	to := *amf.NewTypedObject()
	to.Type = "org.amf.ASClass"
	to.Object = connect.CommandObject

	_, err = encoder.EncodeAmf3Object(buf, to, true)
	if err != nil {
		return nil
	}

	to.Object = connect.OptionalUserArguments
	_, err = encoder.EncodeAmf3Object(buf, to, true)
	if err != nil {
		return nil
	}

	payload = append(payload, buf.Bytes()...)

	return payload
}

// CreateRTMPMessage creates a RTMP message for the connect command message
func (connect *ConnectPayloadAMF3) CreateRTMPMessage() *message.RTMPMessage {
	return &message.RTMPMessage{
		Header:  connect.CreateHeader(),
		Payload: connect.CreatePayload(),
	}
}
