// Package protocol implements the Friends 3DS protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSyncFriend(packet nex.PacketInterface) {
	if protocol.SyncFriend == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends3DS::SyncFriend not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	lfc := types.NewPrimitiveU64(0)
	pids := types.NewList[*types.PID]()
	pids.Type = types.NewPID(0)
	lfcList := types.NewList[*types.PrimitiveU64]()
	lfcList.Type = types.NewPrimitiveU64(0)

	var err error

	err = lfc.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SyncFriend(fmt.Errorf("Failed to read lfc from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = pids.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SyncFriend(fmt.Errorf("Failed to read pids from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lfcList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SyncFriend(fmt.Errorf("Failed to read lfcList from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SyncFriend(nil, packet, callID, lfc, pids, lfcList)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
