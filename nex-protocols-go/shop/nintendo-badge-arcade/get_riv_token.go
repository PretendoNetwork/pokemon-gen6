// Package protocol implements the Nintendo Badge Arcade Shop protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetRivToken(packet nex.PacketInterface) {
	if protocol.GetRivToken == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "ShopNintendoBadgeArcade::GetRivToken not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	itemCode := types.NewString("")
	referenceID := types.NewQBuffer(nil)

	var err error

	err = itemCode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRivToken(fmt.Errorf("Failed to read itemCode from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = referenceID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetRivToken(fmt.Errorf("Failed to read referenceID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetRivToken(nil, packet, callID, itemCode, referenceID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
