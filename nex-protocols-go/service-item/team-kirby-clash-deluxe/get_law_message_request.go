// Package protocol implements the ServiceItemTeamKirbyClashDeluxe protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item_team_kirby_clash_deluxe_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/team-kirby-clash-deluxe/types"
)

func (protocol *Protocol) handleGetLawMessageRequest(packet nex.PacketInterface) {
	if protocol.GetLawMessageRequest == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ServiceItemTeamKirbyClashDeluxe::GetLawMessageRequest not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	getLawMessageParam := service_item_team_kirby_clash_deluxe_types.NewServiceItemGetLawMessageParam()

	err := getLawMessageParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetLawMessageRequest(fmt.Errorf("Failed to read getLawMessageParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetLawMessageRequest(nil, packet, callID, getLawMessageParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
