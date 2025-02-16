// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetDeletionReason(packet nex.PacketInterface) {
	if protocol.SetDeletionReason == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::SetDeletionReason not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	dataIDLst := types.NewList[*types.PrimitiveU64]()
	dataIDLst.Type = types.NewPrimitiveU64(0)
	deletionReason := types.NewPrimitiveU32(0)

	var err error

	err = dataIDLst.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetDeletionReason(fmt.Errorf("Failed to read dataIDLst from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = deletionReason.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetDeletionReason(fmt.Errorf("Failed to read deletionReason from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetDeletionReason(nil, packet, callID, dataIDLst, deletionReason)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
