// Package protocol implements the Persistent Store protocol
package protocol

import (
	"fmt"
	"slices"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the protocol ID for the Persistent Store protocol
	ProtocolID = 0x18

	// MethodFindByGroup is the method ID for method FindByGroup
	MethodFindByGroup = 0x1

	// MethodInsertItem is the method ID for method InsertItem
	MethodInsertItem = 0x2

	// MethodRemoveItem is the method ID for method RemoveItem
	MethodRemoveItem = 0x3

	// MethodGetItem is the method ID for method GetItem
	MethodGetItem = 0x4

	// MethodInsertCustomItem is the method ID for method InsertCustomItem
	MethodInsertCustomItem = 0x5

	// MethodGetCustomItem is the method ID for method GetCustomItem
	MethodGetCustomItem = 0x6

	// MethodFindItemsBySQLQuery is the method ID for method FindItemsBySQLQuery
	MethodFindItemsBySQLQuery = 0x7
)

// Protocol handles the Persistent Store protocol
type Protocol struct {
	endpoint            nex.EndpointInterface
	FindByGroup         func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	InsertItem          func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, bufData *types.Buffer, bReplace *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	RemoveItem          func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error)
	GetItem             func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error)
	InsertCustomItem    func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, hData *types.AnyDataHolder, bReplace *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetCustomItem       func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error)
	FindItemsBySQLQuery func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, strQuery *types.String) (*nex.RMCMessage, *nex.Error)
	Patches             nex.ServiceProtocol
	PatchedMethods      []uint32
}

// Interface implements the methods present on the Persistent Store protocol struct
type Interface interface {
	Endpoint() nex.EndpointInterface
	SetEndpoint(endpoint nex.EndpointInterface)
	SetHandlerFindByGroup(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error))
	SetHandlerInsertItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, bufData *types.Buffer, bReplace *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerRemoveItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerInsertCustomItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, hData *types.AnyDataHolder, bReplace *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error))
	SetHandlerGetCustomItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error))
	SetHandlerFindItemsBySQLQuery(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, strQuery *types.String) (*nex.RMCMessage, *nex.Error))
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
}

// SetHandlerFindByGroup sets the handler for the FindByGroup method
func (protocol *Protocol) SetHandlerFindByGroup(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindByGroup = handler
}

// SetHandlerInsertItem sets the handler for the InsertItem method
func (protocol *Protocol) SetHandlerInsertItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, bufData *types.Buffer, bReplace *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.InsertItem = handler
}

// SetHandlerRemoveItem sets the handler for the RemoveItem method
func (protocol *Protocol) SetHandlerRemoveItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.RemoveItem = handler
}

// SetHandlerGetItem sets the handler for the GetItem method
func (protocol *Protocol) SetHandlerGetItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetItem = handler
}

// SetHandlerInsertCustomItem sets the handler for the InsertCustomItem method
func (protocol *Protocol) SetHandlerInsertCustomItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, hData *types.AnyDataHolder, bReplace *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)) {
	protocol.InsertCustomItem = handler
}

// SetHandlerGetCustomItem sets the handler for the GetCustomItem method
func (protocol *Protocol) SetHandlerGetCustomItem(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.GetCustomItem = handler
}

// SetHandlerFindItemsBySQLQuery sets the handler for the FindItemsBySQLQuery method
func (protocol *Protocol) SetHandlerFindItemsBySQLQuery(handler func(err error, packet nex.PacketInterface, callID uint32, uiGroup *types.PrimitiveU32, strTag *types.String, strQuery *types.String) (*nex.RMCMessage, *nex.Error)) {
	protocol.FindItemsBySQLQuery = handler
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
	case MethodFindByGroup:
		protocol.handleFindByGroup(packet)
	case MethodInsertItem:
		protocol.handleInsertItem(packet)
	case MethodRemoveItem:
		protocol.handleRemoveItem(packet)
	case MethodGetItem:
		protocol.handleGetItem(packet)
	case MethodInsertCustomItem:
		protocol.handleInsertCustomItem(packet)
	case MethodGetCustomItem:
		protocol.handleGetCustomItem(packet)
	case MethodFindItemsBySQLQuery:
		protocol.handleFindItemsBySQLQuery(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Persistent Store method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Persistent Store protocol
func NewProtocol() *Protocol {
	return &Protocol{}
}
