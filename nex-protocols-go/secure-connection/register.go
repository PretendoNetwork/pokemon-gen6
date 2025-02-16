// Package protocol implements the Secure Connection protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRegister(packet nex.PacketInterface) {
	if protocol.Register == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SecureConnection::Register not implemented")

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

	err := vecMyURLs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.Register(fmt.Errorf("Failed to read vecMyURLs from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.Register(nil, packet, callID, vecMyURLs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
