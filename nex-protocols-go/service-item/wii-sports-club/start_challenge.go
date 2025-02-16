// Package protocol implements the ServiceItemWiiSportsClub protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	service_item_wii_sports_club_types "github.com/PretendoNetwork/nex-protocols-go/v2/service-item/wii-sports-club/types"
)

func (protocol *Protocol) handleStartChallenge(packet nex.PacketInterface) {
	if protocol.StartChallenge == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ServiceItemWiiSportsClub::StartChallenge not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	startChallengeParam := service_item_wii_sports_club_types.NewServiceItemStartChallengeParam()

	err := startChallengeParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.StartChallenge(fmt.Errorf("Failed to read startChallengeParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.StartChallenge(nil, packet, callID, startChallengeParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
