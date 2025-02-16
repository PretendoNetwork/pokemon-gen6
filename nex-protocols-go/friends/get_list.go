// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetList(packet nex.PacketInterface) {
	if protocol.GetList == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::GetList not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	byRelationship := types.NewPrimitiveU8(0)
	bReversed := types.NewPrimitiveBool(false)

	var err error

	err = byRelationship.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetList(fmt.Errorf("Failed to read byRelationship from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = bReversed.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetList(fmt.Errorf("Failed to read bReversed from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetList(nil, packet, callID, byRelationship, bReversed)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
