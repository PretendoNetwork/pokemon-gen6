// Package protocol implements the Miitopia DataStore protocol
package protocol

import (
	"fmt"
	"slices"

	"github.com/PretendoNetwork/nex-go/v2"
	datastore "github.com/PretendoNetwork/nex-protocols-go/v2/datastore"
	datastore_miitopia_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/miitopia/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

const (
	// ProtocolID is the Protocol ID for the DataStore (Miitopia) protocol. ID is the same as the DataStore protocol
	ProtocolID = 0x73

	// MethodSearchMii is the method ID for SearchMii
	MethodSearchMii = 0x2F
)

var patchedMethods = []uint32{
	MethodSearchMii,
}

type dataStoreProtocol = datastore.Protocol

// Protocol stores all the RMC method handlers for the DataStore (Miitopia) protocol and listens for requests
// Embeds the DataStore protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	dataStoreProtocol
	SearchMii func(err error, packet nex.PacketInterface, callId uint32, param *datastore_miitopia_types.MiiTubeSearchParam) (*nex.RMCMessage, *nex.Error)
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
	case MethodSearchMii:
		protocol.handleSearchMii(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported DataStoreMiitopia method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new DataStore (Miitopia) protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.dataStoreProtocol.SetEndpoint(endpoint)

	return protocol
}
