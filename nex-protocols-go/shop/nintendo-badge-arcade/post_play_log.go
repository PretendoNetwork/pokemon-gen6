// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	shop_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/v2/shop/nintendo-badge-arcade/types"
)

func (protocol *Protocol) handlePostPlayLog(packet nex.PacketInterface) {
	if protocol.PostPlayLog == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopNintendoBadgeArcade::PostPlayLog not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := shop_nintendo_badge_arcade_types.NewShopPostPlayLogParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PostPlayLog(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PostPlayLog(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
