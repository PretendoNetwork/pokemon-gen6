// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handlePerpetuateObject(packet nex.PacketInterface) {
	if protocol.PerpetuateObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::PerpetuateObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	persistenceSlotID := types.NewPrimitiveU16(0)
	dataID := types.NewPrimitiveU64(0)
	deleteLastObject := types.NewPrimitiveBool(false)

	var err error

	err = persistenceSlotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PerpetuateObject(fmt.Errorf("Failed to read persistenceSlotID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = dataID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PerpetuateObject(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = deleteLastObject.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PerpetuateObject(fmt.Errorf("Failed to read deleteLastObject from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PerpetuateObject(nil, packet, callID, persistenceSlotID, dataID, deleteLastObject)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
