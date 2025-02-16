// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-bank/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handlePrepareUpdateBankObject(packet nex.PacketInterface) {
	if protocol.PrepareUpdateBankObject == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonBank::PrepareUpdateBankObject not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	transactionParam := datastore_pokemon_bank_types.NewBankTransactionParam()

	err := transactionParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PrepareUpdateBankObject(fmt.Errorf("Failed to read transactionParam from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PrepareUpdateBankObject(nil, packet, callID, transactionParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
