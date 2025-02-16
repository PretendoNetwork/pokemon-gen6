// Package protocol implements the DataStorePokemonGen6 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_pokemon_gen6_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-gen6/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Pokemon Gen6) protocol
	ProtocolID = 0x73

	// MethodPrepareUploadPokemon is the method ID for the PrepareUploadPokemon method
	MethodPrepareUploadPokemon = 0x28

	// MethodUploadPokemon is the method ID for the UploadPokemon method
	MethodUploadPokemon = 0x29

	// MethodSearchPokemon is the method ID for the SearchPokemon method
	MethodSearchPokemon = 0x2A

	// MethodPrepareTradePokemon is the method ID for the PrepareTradePokemon method
	MethodPrepareTradePokemon = 0x2B

	// MethodTradePokemon is the method ID for the TradePokemon method
	MethodTradePokemon = 0x2C

	// MethodDownloadOtherPokemon is the method ID for the DownloadOtherPokemon method
	MethodDownloadOtherPokemon = 0x2D

	// MethodDownloadMyPokemon is the method ID for the DownloadMyPokemon method
	MethodDownloadMyPokemon = 0x2E

	// MethodDeletePokemon is the method ID for the DeletePokemon method
	MethodDeletePokemon = 0x2F
)

var patchedMethods = []uint32{
	MethodPrepareUploadPokemon,
	MethodUploadPokemon,
	MethodSearchPokemon,
	MethodPrepareTradePokemon,
	MethodTradePokemon,
	MethodDownloadOtherPokemon,
	MethodDownloadMyPokemon,
	MethodDeletePokemon,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Pokemon Gen6) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	PrepareUploadPokemon func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	UploadPokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationUploadPokemonParam) (*nex.RMCMessage, *nex.Error)
	SearchPokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationSearchPokemonParam) (*nex.RMCMessage, *nex.Error)
	PrepareTradePokemon  func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationPrepareTradePokemonParam) (*nex.RMCMessage, *nex.Error)
	TradePokemon         func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationTradePokemonParam) (*nex.RMCMessage, *nex.Error)
	DownloadOtherPokemon func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadOtherPokemonParam) (*nex.RMCMessage, *nex.Error)
	DownloadMyPokemon    func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDownloadMyPokemonParam) (*nex.RMCMessage, *nex.Error)
	DeletePokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_gen6_types.GlobalTradeStationDeletePokemonParam) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.dataStoreProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodPrepareUploadPokemon:
		protocol.handlePrepareUploadPokemon(packet)
	case MethodUploadPokemon:
		protocol.handleUploadPokemon(packet)
	case MethodSearchPokemon:
		protocol.handleSearchPokemon(packet)
	case MethodPrepareTradePokemon:
		protocol.handlePrepareTradePokemon(packet)
	case MethodTradePokemon:
		protocol.handleTradePokemon(packet)
	case MethodDownloadOtherPokemon:
		protocol.handleDownloadOtherPokemon(packet)
	case MethodDownloadMyPokemon:
		protocol.handleDownloadMyPokemon(packet)
	case MethodDeletePokemon:
		protocol.handleDeletePokemon(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported DataStore (Pokemon Gen6) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new DataStorePokemonGen6 protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.dataStoreProtocol.SetEndpoint(endpoint)

	return protocol
}
