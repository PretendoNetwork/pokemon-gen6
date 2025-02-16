// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetPersistenceInfos(packet nex.PacketInterface) {
	if protocol.GetPersistenceInfos == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::GetPersistenceInfos not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ownerID := types.NewPID(0)
	persistenceSlotIDs := types.NewList[*types.PrimitiveU16]()
	persistenceSlotIDs.Type = types.NewPrimitiveU16(0)

	var err error

	err = ownerID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPersistenceInfos(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = persistenceSlotIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetPersistenceInfos(fmt.Errorf("Failed to read persistenceSlotIDs from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetPersistenceInfos(nil, packet, callID, ownerID, persistenceSlotIDs)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
