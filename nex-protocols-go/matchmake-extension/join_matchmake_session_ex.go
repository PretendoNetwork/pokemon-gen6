// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleJoinMatchmakeSessionEx(packet nex.PacketInterface) {
	if protocol.JoinMatchmakeSessionEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::JoinMatchmakeSessionEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	matchmakingVersion := endpoint.LibraryVersions().MatchMaking
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	gid := types.NewPrimitiveU32(0)
	strMessage := types.NewString("")
	dontCareMyBlockList := types.NewPrimitiveBool(false)
	participationCount := types.NewPrimitiveU16(0)

	var err error

	err = gid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read gid from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strMessage.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read strMessage from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = dontCareMyBlockList.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read dontCareMyBlockList from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	if matchmakingVersion.GreaterOrEqual("3.4.0") {
		err = participationCount.ExtractFrom(parametersStream)
		if err != nil {
			_, rmcError := protocol.JoinMatchmakeSessionEx(fmt.Errorf("Failed to read participationCount from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}
	}

	rmcMessage, rmcError := protocol.JoinMatchmakeSessionEx(nil, packet, callID, gid, strMessage, dontCareMyBlockList, participationCount)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
