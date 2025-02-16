// Package protocol implements the Shop protocol
package protocol

// * Stubbed, Kinnay documents this as being game-specific for Pokemon bank however Badge Arcade and Pokemon gen 7 uses this protocol as well
// TODO - Figure out more about this protocol, unsure if anything here is right

import (
	"fmt"
	"slices"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the Protocol ID for the Shop protocol
	ProtocolID = 0xC8
)

// Protocol stores all the RMC method handlers for the Shop protocol and listens for requests
type Protocol struct {
	endpoint       nex.EndpointInterface
	Patches        nex.ServiceProtocol
	PatchedMethods []uint32
}

// Endpoint returns the endpoint implementing the protocol
func (protocol *Protocol) Endpoint() nex.EndpointInterface {
	return protocol.endpoint
}

// SetEndpoint sets the endpoint implementing the protocol
func (protocol *Protocol) SetEndpoint(endpoint nex.EndpointInterface) {
	protocol.endpoint = endpoint
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
	default:
		errMessage := fmt.Sprintf("Shop method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}
