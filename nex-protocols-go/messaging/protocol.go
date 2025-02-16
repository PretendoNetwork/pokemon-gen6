// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
)

const (
	// ProtocolID is the protocol ID for the Messaging protocol
	ProtocolID = 0x17

	// MethodDeliverMessage is the method ID for method DeliverMessage
	MethodDeliverMessage = 0x1

	// MethodGetNumberOfMessages is the method ID for method GetNumberOfMessages
	MethodGetNumberOfMessages = 0x2

	// MethodGetMessagesHeaders is the method ID for method GetMessagesHeaders
	MethodGetMessagesHeaders = 0x3

	// MethodRetrieveAllMessagesWithinRange is the method ID for method RetrieveAllMessagesWithinRange
	MethodRetrieveAllMessagesWithinRange = 0x4

	// MethodRetrieveMessages is the method ID for method RetrieveMessages
	MethodRetrieveMessages = 0x5

	// MethodDeleteMessages is the method ID for method DeleteMessages
	MethodDeleteMessages = 0x6

	// MethodDeleteAllMessages is the method ID for method DeleteAllMessages
	MethodDeleteAllMessages = 0x7

	// MethodDeliverMessageMultiTarget is the method ID for method DeliverMessageMultiTarget
	MethodDeliverMessageMultiTarget = 0x8
)

// Protocol stores all the RMC method handlers for the Messaging protocol and listens for requests
type Protocol struct {
	endpoint                       nex.EndpointInterface
	DeliverMessage                 func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)
	GetNumberOfMessages            func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, *nex.Error)
	GetMessagesHeaders             func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	RetrieveAllMessagesWithinRange func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)
	RetrieveMessages               func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs *types.List[*types.PrimitiveU32], bLeaveOnServer *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	DeleteMessages                 func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
	DeleteAllMessages              func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, *nex.Error)
	DeliverMessageMultiTarget      func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error) // TODO - Unknown request/response format
	Patches                        nex.ServiceProtocol
	PatchedMethods                 []uint32
}

// Interface implements the methods present on the Messaging protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerDeliverMessage(handler func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetNumberOfMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetMessagesHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerRetrieveAllMessagesWithinRange(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error))
	SetHandlerRetrieveMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs *types.List[*types.PrimitiveU32], bLeaveOnServer *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeleteAllMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, *nex.Error))
	SetHandlerDeliverMessageMultiTarget(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerDeliverMessage sets the handler for the DeliverMessage method
func (protocol *Protocol) SetHandlerDeliverMessage(handler func(err error, packet nex.PacketInterface, callID uint32, oUserMessage *types.AnyDataHolder) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeliverMessage = handler
}

// SetHandlerGetNumberOfMessages sets the handler for the GetNumberOfMessages method
func (protocol *Protocol) SetHandlerGetNumberOfMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetNumberOfMessages = handler
}

// SetHandlerGetMessagesHeaders sets the handler for the GetMessagesHeaders method
func (protocol *Protocol) SetHandlerGetMessagesHeaders(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetMessagesHeaders = handler
}

// SetHandlerRetrieveAllMessagesWithinRange sets the handler for the RetrieveAllMessagesWithinRange method
func (protocol *Protocol) SetHandlerRetrieveAllMessagesWithinRange(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, resultRange *types.ResultRange) (*nex.RMCMessage, *nex.Error)) {
	protocol.RetrieveAllMessagesWithinRange = handler
}

// SetHandlerRetrieveMessages sets the handler for the RetrieveMessages method
func (protocol *Protocol) SetHandlerRetrieveMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMsgIDs *types.List[*types.PrimitiveU32], bLeaveOnServer *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.RetrieveMessages = handler
}

// SetHandlerDeleteMessages sets the handler for the DeleteMessages method
func (protocol *Protocol) SetHandlerDeleteMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient, lstMessagesToDelete *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteMessages = handler
}

// SetHandlerDeleteAllMessages sets the handler for the DeleteAllMessages method
func (protocol *Protocol) SetHandlerDeleteAllMessages(handler func(err error, packet nex.PacketInterface, callID uint32, recipient *messaging_types.MessageRecipient) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeleteAllMessages = handler
}

// SetHandlerDeliverMessageMultiTarget sets the handler for the DeliverMessageMultiTarget method
func (protocol *Protocol) SetHandlerDeliverMessageMultiTarget(handler func(err error, packet nex.PacketInterface, callID uint32, packetPayload []byte) (*nex.RMCMessage, *nex.Error)) {
	protocol.DeliverMessageMultiTarget = handler
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if protocol.Patches != nil && slices.Contains(protocol.PatchedMethods, message.MethodID) {
		protocol.Patches.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodDeliverMessage:
		protocol.handleDeliverMessage(packet)
	case MethodGetNumberOfMessages:
		protocol.handleGetNumberOfMessages(packet)
	case MethodGetMessagesHeaders:
		protocol.handleGetMessagesHeaders(packet)
	case MethodRetrieveAllMessagesWithinRange:
		protocol.handleRetrieveAllMessagesWithinRange(packet)
	case MethodRetrieveMessages:
		protocol.handleRetrieveMessages(packet)
	case MethodDeleteMessages:
		protocol.handleDeleteMessages(packet)
	case MethodDeleteAllMessages:
		protocol.handleDeleteAllMessages(packet)
	case MethodDeliverMessageMultiTarget:
		protocol.handleDeliverMessageMultiTarget(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Messaging method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Messaging protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
