// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleDownloadMyPokemon(packet nex.PacketInterface) {
	if protocol.DownloadMyPokemon == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStorePokemonGen6::DownloadMyPokemon not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := datastore_pokemon_gen6_types.NewGlobalTradeStationDownloadMyPokemonParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.DownloadMyPokemon(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.DownloadMyPokemon(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
