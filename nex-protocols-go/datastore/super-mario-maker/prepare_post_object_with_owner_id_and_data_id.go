// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handlePreparePostObjectWithOwnerIDAndDataID(packet nex.PacketInterface) {
	if protocol.PreparePostObjectWithOwnerIDAndDataID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::PreparePostObjectWithOwnerIDAndDataID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	ownerID := types.NewPrimitiveU32(0)
	dataID := types.NewPrimitiveU64(0)
	param := datastore_types.NewDataStorePreparePostParam()

	var err error

	err = ownerID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read ownerID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = dataID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read dataID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PreparePostObjectWithOwnerIDAndDataID(nil, packet, callID, ownerID, dataID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
