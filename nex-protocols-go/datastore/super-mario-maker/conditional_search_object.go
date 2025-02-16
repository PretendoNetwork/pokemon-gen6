// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleConditionalSearchObject(packet nex.PacketInterface) {
	if protocol.ConditionalSearchObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::ConditionalSearchObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	condition := types.NewPrimitiveU32(0)
	param := datastore_types.NewDataStoreSearchParam()
	extraData := types.NewList[*types.String]()
	extraData.Type = types.NewString("")

	var err error

	err = condition.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ConditionalSearchObject(fmt.Errorf("Failed to read condition from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ConditionalSearchObject(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = extraData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ConditionalSearchObject(fmt.Errorf("Failed to read extraData from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ConditionalSearchObject(nil, packet, callID, condition, param, extraData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
