// Package protocol implements the DataStorePokemonBank protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_pokemon_bank_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/pokemon-bank/types"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Pokemon Bank) protocol
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

	// MethodGetTransactionParam is the method ID for the GetTransactionParam method
	MethodGetTransactionParam = 0x30

	// MethodPreparePostBankObject is the method ID for the PreparePostBankObject method
	MethodPreparePostBankObject = 0x31

	// MethodCompletePostBankObject is the method ID for the CompletePostBankObject method
	MethodCompletePostBankObject = 0x32

	// MethodPrepareGetBankObject is the method ID for the PrepareGetBankObject method
	MethodPrepareGetBankObject = 0x33

	// MethodPrepareUpdateBankObject is the method ID for the PrepareUpdateBankObject method
	MethodPrepareUpdateBankObject = 0x34

	// MethodCompleteUpdateBankObject is the method ID for the CompleteUpdateBankObject method
	MethodCompleteUpdateBankObject = 0x35

	// MethodRollbackBankObject is the method ID for the RollbackBankObject method
	MethodRollbackBankObject = 0x36

	// MethodGetUnlockKey is the method ID for the GetUnlockKey method
	MethodGetUnlockKey = 0x37

	// MethodRequestMigration is the method ID for the RequestMigration method
	MethodRequestMigration = 0x38
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
	MethodGetTransactionParam,
	MethodPreparePostBankObject,
	MethodCompletePostBankObject,
	MethodPrepareGetBankObject,
	MethodPrepareUpdateBankObject,
	MethodCompleteUpdateBankObject,
	MethodRollbackBankObject,
	MethodGetUnlockKey,
	MethodRequestMigration,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Pokemon Bank) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	PrepareUploadPokemon     func(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error)
	UploadPokemon            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationUploadPokemonParam) (*nex.RMCMessage, *nex.Error)
	SearchPokemon            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationSearchPokemonParam) (*nex.RMCMessage, *nex.Error)
	PrepareTradePokemon      func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationPrepareTradePokemonParam) (*nex.RMCMessage, *nex.Error)
	TradePokemon             func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationTradePokemonParam) (*nex.RMCMessage, *nex.Error)
	DownloadOtherPokemon     func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDownloadOtherPokemonParam) (*nex.RMCMessage, *nex.Error)
	DownloadMyPokemon        func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDownloadMyPokemonParam) (*nex.RMCMessage, *nex.Error)
	DeletePokemon            func(err error, packet nex.PacketInterface, callID uint32, param *datastore_pokemon_bank_types.GlobalTradeStationDeletePokemonParam) (*nex.RMCMessage, *nex.Error)
	GetTransactionParam      func(err error, packet nex.PacketInterface, callID uint32, slotID *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)
	PreparePostBankObject    func(err error, packet nex.PacketInterface, callID uint32, slotID *types.PrimitiveU16, size *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	CompletePostBankObject   func(err error, packet nex.PacketInterface, callID uint32, param *datastore_types.DataStoreCompletePostParam) (*nex.RMCMessage, *nex.Error)
	PrepareGetBankObject     func(err error, packet nex.PacketInterface, callID uint32, slotID *types.PrimitiveU16, applicationID *types.PrimitiveU16) (*nex.RMCMessage, *nex.Error)
	PrepareUpdateBankObject  func(err error, packet nex.PacketInterface, callID uint32, transactionParam *datastore_pokemon_bank_types.BankTransactionParam) (*nex.RMCMessage, *nex.Error)
	CompleteUpdateBankObject func(err error, packet nex.PacketInterface, callID uint32, slotID *types.PrimitiveU16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	RollbackBankObject       func(err error, packet nex.PacketInterface, callID uint32, slotID *types.PrimitiveU16, transactionParam *datastore_pokemon_bank_types.BankTransactionParam, isForce *types.PrimitiveBool) (*nex.RMCMessage, *nex.Error)
	GetUnlockKey             func(err error, packet nex.PacketInterface, callID uint32, challengeValue *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	RequestMigration         func(err error, packet nex.PacketInterface, callID uint32, oneTimePassword *types.String, boxes *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
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
	case MethodGetTransactionParam:
		protocol.handleGetTransactionParam(packet)
	case MethodPreparePostBankObject:
		protocol.handlePreparePostBankObject(packet)
	case MethodCompletePostBankObject:
		protocol.handleCompletePostBankObject(packet)
	case MethodPrepareGetBankObject:
		protocol.handlePrepareGetBankObject(packet)
	case MethodPrepareUpdateBankObject:
		protocol.handlePrepareUpdateBankObject(packet)
	case MethodCompleteUpdateBankObject:
		protocol.handleCompleteUpdateBankObject(packet)
	case MethodRollbackBankObject:
		protocol.handleRollbackBankObject(packet)
	case MethodGetUnlockKey:
		protocol.handleGetUnlockKey(packet)
	case MethodRequestMigration:
		protocol.handleRequestMigration(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported DataStore (Pokemon Bank) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new DataStorePokemonBank protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.dataStoreProtocol.SetEndpoint(endpoint)

	return protocol
}
