// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetServiceItemRightRequest(packet nex.PacketInterface) {
	if protocol.GetServiceItemRightRequest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ServiceItemTeamKirbyClashDeluxe::GetServiceItemRightRequest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	getServiceItemRightParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemGetServiceItemRightParam()
	withoutRightBinary := types.NewPrimitiveBool(false)

	var err error

	err = getServiceItemRightParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetServiceItemRightRequest(fmt.Errorf("Failed to read getServiceItemRightParam from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = withoutRightBinary.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetServiceItemRightRequest(fmt.Errorf("Failed to read withoutRightBinary from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetServiceItemRightRequest(nil, packet, callID, getServiceItemRightParam, withoutRightBinary)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
