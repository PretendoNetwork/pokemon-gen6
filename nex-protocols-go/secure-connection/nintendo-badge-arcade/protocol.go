// Package protocol implements the Nintendo Badge Arcade Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	secure_connection "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the Secure Connection (Nintendo Badge Arcade) protocol. ID is the same as the Secure Connection Protocol
	ProtocolID = 0xB

	// MethodGetMaintenanceStatus is the method ID for GetMaintenanceStatus
	MethodGetMaintenanceStatus = 0x9
)

var patchedMethods = []uint32{
	MethodGetMaintenanceStatus,
}

type secureConnectionProtocol = secure_connection.Protocol

// Protocol stores all the RMC method handlers for the Secure Connection (Nintendo Badge Arcade) protocol and listens for requests
// Embeds the SecureConnection Protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	secureConnectionProtocol
	GetMaintenanceStatus func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.secureConnectionProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetMaintenanceStatus:
		protocol.handleGetMaintenanceStatus(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported SecureConnectionNintendoBadgeArcade method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Secure Connection (Nintendo Badge Arcade) protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.secureConnectionProtocol.SetEndpoint(endpoint)

	return protocol
}
