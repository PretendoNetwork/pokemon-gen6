// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

func (protocol *Protocol) handleBrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(packet nex.PacketInterface) {
	if protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	searchCriteria := match_making_types.NewMatchmakeSessionSearchCriteria()

	err := searchCriteria.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(fmt.Errorf("Failed to read searchCriteria from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BrowseMatchmakeSessionWithHostURLsNoHolderNoResultRange(nil, packet, callID, searchCriteria)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
