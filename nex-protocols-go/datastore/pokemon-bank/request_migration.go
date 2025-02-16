// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRequestMigration(packet nex.PacketInterface) {
	if protocol.RequestMigration == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::RequestMigration not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	oneTimePassword := types.NewString("")
	boxes := types.NewList[*types.PrimitiveU32]()
	boxes.Type = types.NewPrimitiveU32(0)

	var err error

	err = oneTimePassword.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestMigration(fmt.Errorf("Failed to read oneTimePassword from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = boxes.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestMigration(fmt.Errorf("Failed to read boxes from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RequestMigration(nil, packet, callID, oneTimePassword, boxes)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
