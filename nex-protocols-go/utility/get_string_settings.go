// Package protocol implements the Utility protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetStringSettings(packet nex.PacketInterface) {
	if protocol.GetStringSettings == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Utility::GetStringSettings not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	stringSettingIndex := types.NewPrimitiveU32(0)

	err := stringSettingIndex.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetStringSettings(fmt.Errorf("Failed to read stringSettingIndex from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetStringSettings(nil, packet, callID, stringSettingIndex)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
