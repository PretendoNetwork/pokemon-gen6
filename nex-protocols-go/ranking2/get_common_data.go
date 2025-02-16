// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetCommonData(packet nex.PacketInterface) {
	if protocol.GetCommonData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::GetCommonData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	optionFlags := types.NewPrimitiveU32(0)
	principalID := types.NewPID(0)
	nexUniqueID := types.NewPrimitiveU64(0)

	var err error

	err = optionFlags.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCommonData(fmt.Errorf("Failed to read optionFlags from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = principalID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCommonData(fmt.Errorf("Failed to read principalID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = nexUniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetCommonData(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetCommonData(nil, packet, callID, optionFlags, principalID, nexUniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
