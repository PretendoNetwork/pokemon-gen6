// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handlePurchaseServiceItemRequest(packet nex.PacketInterface) {
	if protocol.PurchaseServiceItemRequest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ServiceItemWiiSportsClub::PurchaseServiceItemRequest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	purchaseServiceItemParam := service_item_wii_sports_club_types.NewServiceItemPurchaseServiceItemParam()

	err := purchaseServiceItemParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PurchaseServiceItemRequest(fmt.Errorf("Failed to read purchaseServiceItemParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PurchaseServiceItemRequest(nil, packet, callID, purchaseServiceItemParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
