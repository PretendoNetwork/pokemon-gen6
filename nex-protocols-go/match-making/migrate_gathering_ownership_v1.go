// Package protocol implements the Match Making protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleMigrateGatheringOwnershipV1(packet nex.PacketInterface) {
	if protocol.MigrateGatheringOwnershipV1 == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchMaking::MigrateGatheringOwnershipV1 not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	gid := types.NewPrimitiveU32(0)
	lstPotentialNewOwnersID := types.NewList[*types.PID]()
	lstPotentialNewOwnersID.Type = types.NewPID(0)

	var err error

	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.MigrateGatheringOwnershipV1(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstPotentialNewOwnersID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.MigrateGatheringOwnershipV1(fmt.Errorf("Failed to read lstPotentialNewOwnersID from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.MigrateGatheringOwnershipV1(nil, packet, callID, gid, lstPotentialNewOwnersID)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
