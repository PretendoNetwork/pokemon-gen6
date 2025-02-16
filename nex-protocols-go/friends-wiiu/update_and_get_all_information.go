// Package protocol implements the Friends WiiU protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	friends_wiiu_types "github.com/PretendoNetwork/nex-protocols-go/v2/friends-wiiu/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateAndGetAllInformation(packet nex.PacketInterface) {
	if protocol.UpdateAndGetAllInformation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "FriendsWiiU::UpdateAndGetAllInformation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	nnaInfo := friends_wiiu_types.NewNNAInfo()
	presence := friends_wiiu_types.NewNintendoPresenceV2()
	birthday := types.NewDateTime(0)

	var err error

	err = nnaInfo.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read nnaInfo from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = presence.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read presence from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = birthday.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateAndGetAllInformation(fmt.Errorf("Failed to read birthday from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateAndGetAllInformation(nil, packet, callID, nnaInfo, presence, birthday)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
