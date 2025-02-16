// Package protocol implements the AAUser protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	aauser_types "github.com/PretendoNetwork/nex-protocols-go/v2/aa-user/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetApplicationInfo(packet nex.PacketInterface) {
	if protocol.SetApplicationInfo == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AAUser::SetApplicationInfo not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	applicationInfo := types.NewList[*aauser_types.ApplicationInfo]()
	applicationInfo.Type = aauser_types.NewApplicationInfo()

	err := applicationInfo.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetApplicationInfo(fmt.Errorf("Failed to read applicationInfo from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetApplicationInfo(nil, packet, callID, applicationInfo)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
