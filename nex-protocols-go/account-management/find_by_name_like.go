// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleFindByNameLike(packet nex.PacketInterface) {
	if protocol.FindByNameLike == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::FindByNameLike not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	uiGroups := types.NewPrimitiveU32(0)
	strLike := types.NewString("")
	resultRange := types.NewResultRange()

	var err error

	err = uiGroups.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameLike(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strLike.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameLike(fmt.Errorf("Failed to read strLike from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = resultRange.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.FindByNameLike(fmt.Errorf("Failed to read resultRange from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.FindByNameLike(nil, packet, callID, uiGroups, strLike, resultRange)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
