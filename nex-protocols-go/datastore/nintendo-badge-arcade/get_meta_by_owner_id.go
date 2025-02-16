// Package protocol implements the Nintendo Badge Arcade DataStore protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	datastore_nintendo_badge_arcade_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/nintendo-badge-arcade/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleGetMetaByOwnerID(packet nex.PacketInterface) {
	if protocol.GetMetaByOwnerID == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreBadgeArcade::GetMetaByOwnerID not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	param := datastore_nintendo_badge_arcade_types.NewDataStoreGetMetaByOwnerIDParam()

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.GetMetaByOwnerID(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.GetMetaByOwnerID(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
