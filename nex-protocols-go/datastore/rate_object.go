// Package protocol implements the DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRateObject(packet nex.PacketInterface) {
	if protocol.RateObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStore::RateObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	target := datastore_types.NewDataStoreRatingTarget()
	param := datastore_types.NewDataStoreRateObjectParam()
	fetchRatings := types.NewPrimitiveBool(false)

	var err error

	err = target.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObject(fmt.Errorf("Failed to read target from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObject(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = fetchRatings.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RateObject(fmt.Errorf("Failed to read fetchRatings from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RateObject(nil, packet, callID, target, param, fetchRatings)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
