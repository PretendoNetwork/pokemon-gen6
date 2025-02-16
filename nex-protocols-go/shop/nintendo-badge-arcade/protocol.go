// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	shop "github.com/PretendoNetwork/nex-protocols-go/v2/shop"
	shop_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/v2/shop/nintendo-badge-arcade/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Shop (Nintendo Badge Arcade) protocol
	ProtocolID = 0xC8

	// MethodGetRivToken is the method ID for GetRivToken
	MethodGetRivToken = 0x1

	// MethodPostPlayLog is the method ID for PostPlayLog
	MethodPostPlayLog = 0x2
)

var patchedMethods = []uint32{
	MethodGetRivToken,
	MethodPostPlayLog,
}

type shopProtocol = shop.Protocol

// Protocol stores all the RMC method handlers for the Shop (Nintendo Badge Arcade) protocol and listens for requests
// Embeds the Shop protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	shopProtocol
	GetRivToken func(err error, packet nex.PacketInterface, callID uint32, itemCode *types.String, referenceID *types.QBuffer) (*nex.RMCMessage, *nex.Error)
	PostPlayLog func(err error, packet nex.PacketInterface, callID uint32, param *shop_nintendo_badge_arcade_types.ShopPostPlayLogParam) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.shopProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodGetRivToken:
		protocol.handleGetRivToken(packet)
	case MethodPostPlayLog:
		protocol.handlePostPlayLog(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported ShopNintendoBadgeArcade method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new Shop (Nintendo Badge Arcade)
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.shopProtocol.SetEndpoint(endpoint)

	return protocol
}
