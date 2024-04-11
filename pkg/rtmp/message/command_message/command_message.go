package command_message

import (
	amf "github.com/elobuff/goamf"
)

// CommandMessagePayload is the payload of the command message
type CommandMessagePayload struct {
	CommandName   string
	TransactionID uint
	CommandObject amf.Object
}
