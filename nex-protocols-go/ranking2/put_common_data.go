// Package protocol implements the Ranking2 protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	ranking2_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking2/types"
)

func (protocol *Protocol) handlePutCommonData(packet nex.PacketInterface) {
	if protocol.PutCommonData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Ranking2::PutCommonData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	commonData := ranking2_types.NewRanking2CommonData()
	nexUniqueID := types.NewPrimitiveU64(0)

	var err error

	err = commonData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PutCommonData(fmt.Errorf("Failed to read commonData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = nexUniqueID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.PutCommonData(fmt.Errorf("Failed to read nexUniqueID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.PutCommonData(nil, packet, callID, commonData, nexUniqueID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
