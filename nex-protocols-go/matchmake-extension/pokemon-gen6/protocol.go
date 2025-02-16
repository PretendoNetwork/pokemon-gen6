// Package protocol implements the Pokemon GEN 6 Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the protocol ID for the Matchmake Extension (Pokemon GEN 6) protocol. ID is the same as the Matchmake Extension protocol
	ProtocolID = 0x6D

	// MethodClearMyPreviouslyMatchedUserCache is the method ID for the method ClearMyPreviouslyMatchedUserCache
	MethodClearMyPreviouslyMatchedUserCache = 0x22
)

var patchedMethods = []uint32{
	MethodClearMyPreviouslyMatchedUserCache,
}

type matchmakeExtensionProtocol = matchmake_extension.Protocol

// Protocol stores all the RMC method handlers for the Matchmake Extension (Pokemon GEN 6) protocol and listens for requests
// Embeds the Matchmake Extension protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	matchmakeExtensionProtocol
	ClearMyPreviouslyMatchedUserCache func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.matchmakeExtensionProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodClearMyPreviouslyMatchedUserCache:
		protocol.handleClearMyPreviouslyMatchedUserCache(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported MatchmakeExtension (Pokemon GEN 6) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Matchmake Extension (Pokemon GEN 6) protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.matchmakeExtensionProtocol.SetEndpoint(endpoint)

	return protocol
}
