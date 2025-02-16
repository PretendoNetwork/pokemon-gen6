// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleCompleteUpdateBankObject(packet nex.PacketInterface) {
	if protocol.CompleteUpdateBankObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::CompleteUpdateBankObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	slotID := types.NewPrimitiveU16(0)
	transactionParam := datastore_pokemon_bank_types.NewBankTransactionParam()
	isForce := types.NewPrimitiveBool(false)

	var err error

	err = slotID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CompleteUpdateBankObject(fmt.Errorf("Failed to read slotID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = transactionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CompleteUpdateBankObject(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = isForce.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CompleteUpdateBankObject(fmt.Errorf("Failed to read isForce from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CompleteUpdateBankObject(nil, packet, callID, slotID, transactionParam, isForce)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
