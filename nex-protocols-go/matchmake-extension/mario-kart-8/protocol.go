// Package protocol implements the MatchmakeExtensionMarioKart8 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	matchmake_extension_mario_kart8_types "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension/mario-kart-8/types"
	"golang.org/x/exp/slices"
)

const (
	// ProtocolID is the Protocol ID for the Matchmake Extension (Mario Kart 8) protocol
	ProtocolID = 0x6D

	// MethodCreateSimpleSearchObject is the method ID for the CreateSimpleSearchObject method
	MethodCreateSimpleSearchObject = 0x24

	// MethodUpdateSimpleSearchObject is the method ID for the UpdateSimpleSearchObject method
	MethodUpdateSimpleSearchObject = 0x25

	// MethodDeleteSimpleSearchObject is the method ID for the DeleteSimpleSearchObject method
	MethodDeleteSimpleSearchObject = 0x26

	// MethodSearchSimpleSearchObject is the method ID for the SearchSimpleSearchObject method
	MethodSearchSimpleSearchObject = 0x27

	// MethodJoinMatchmakeSessionWithExtraParticipants is the method ID for the JoinMatchmakeSessionWithExtraParticipants method
	MethodJoinMatchmakeSessionWithExtraParticipants = 0x28

	// MethodSearchSimpleSearchObjectByObjectIDs is the method ID for the SearchSimpleSearchObjectByObjectIDs method
	MethodSearchSimpleSearchObjectByObjectIDs = 0x29
)

var patchedMethods = []uint32{
	MethodCreateSimpleSearchObject,
	MethodUpdateSimpleSearchObject,
	MethodDeleteSimpleSearchObject,
	MethodSearchSimpleSearchObject,
	MethodJoinMatchmakeSessionWithExtraParticipants,
	MethodSearchSimpleSearchObjectByObjectIDs,
}

type matchmakeExtensionProtocol = matchmake_extension.Protocol

// Protocol stores all the RMC method handlers for the Matchmake Extension (Mario Kart 8) protocol and listens for requests
// Embeds the Matchmake Extension protocol
type Protocol struct {
	endpoint nex.EndpointInterface
	matchmakeExtensionProtocol
	CreateSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, object *matchmake_extension_mario_kart8_types.SimpleSearchObject) (*nex.RMCMessage, *nex.Error)
	UpdateSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, objectID *types.PrimitiveU32, newObject *matchmake_extension_mario_kart8_types.SimpleSearchObject) (*nex.RMCMessage, *nex.Error)
	DeleteSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, objectID *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	SearchSimpleSearchObject                  func(err error, packet nex.PacketInterface, callID uint32, param *matchmake_extension_mario_kart8_types.SimpleSearchParam) (*nex.RMCMessage, *nex.Error)
	JoinMatchmakeSessionWithExtraParticipants func(err error, packet nex.PacketInterface, callID uint32, gid *types.PrimitiveU32, joinMessage *types.String, ignoreBlacklist *types.PrimitiveBool, participationCount *types.PrimitiveU16, extraParticipants *types.PrimitiveU32) (*nex.RMCMessage, *nex.Error)
	SearchSimpleSearchObjectByObjectIDs       func(err error, packet nex.PacketInterface, callID uint32, objectIDs *types.List[*types.PrimitiveU32]) (*nex.RMCMessage, *nex.Error)
}

// HandlePacket sends the packet to the correct RMC method handler
func (protocol *Protocol) HandlePacket(packet nex.PacketInterface) {
	message := packet.RMCMessage()

	if !message.IsRequest || message.ProtocolID != ProtocolID {
		return
	}

	if !slices.Contains(patchedMethods, message.MethodID) {
		protocol.matchmakeExtensionProtocol.HandlePacket(packet)
		return
	}

	switch message.MethodID {
	case MethodCreateSimpleSearchObject:
		protocol.handleCreateSimpleSearchObject(packet)
	case MethodUpdateSimpleSearchObject:
		protocol.handleUpdateSimpleSearchObject(packet)
	case MethodDeleteSimpleSearchObject:
		protocol.handleDeleteSimpleSearchObject(packet)
	case MethodSearchSimpleSearchObject:
		protocol.handleSearchSimpleSearchObject(packet)
	case MethodJoinMatchmakeSessionWithExtraParticipants:
		protocol.handleJoinMatchmakeSessionWithExtraParticipants(packet)
	case MethodSearchSimpleSearchObjectByObjectIDs:
		protocol.handleSearchSimpleSearchObjectByObjectIDs(packet)
	default:
		errMessage := fmt.Sprintf("Unsupported Matchmake Extension (Mario Kart 8) method ID: %#v\n", message.MethodID)
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, errMessage)

		globals.RespondError(packet, ProtocolID, err)
		globals.Logger.Warning(err.Message)
	}
}

// NewProtocol returns a new MatchmakeExtensionMarioKart8 protocol
func NewProtocol(endpoint nex.EndpointInterface) *Protocol {
	protocol := &Protocol{endpoint: endpoint}
	protocol.matchmakeExtensionProtocol.SetEndpoint(endpoint)

	return protocol
}
