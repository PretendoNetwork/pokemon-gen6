// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRegisterEx(packet nex.PacketInterface) {
	if protocol.RegisterEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SecureConnection::RegisterEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	vecMyURLs := types.NewList[*types.StationURL]()
	vecMyURLs.Type = types.NewStationURL("")
	hCustomData := types.NewAnyDataHolder()

	var err error

	err = vecMyURLs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RegisterEx(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = hCustomData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RegisterEx(fmt.Errorf("Failed to read hCustomData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RegisterEx(nil, packet, callID, vecMyURLs, hCustomData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
