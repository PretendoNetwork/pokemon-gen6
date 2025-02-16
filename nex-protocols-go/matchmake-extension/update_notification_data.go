// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleUpdateNotificationData(packet nex.PacketInterface) {
	if protocol.UpdateNotificationData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::UpdateNotificationData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	uiType := types.NewPrimitiveU32(0)
	uiParam1 := types.NewPrimitiveU32(0)
	uiParam2 := types.NewPrimitiveU32(0)
	strParam := types.NewString("")

	var err error

	err = uiType.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiType from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uiParam1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiParam1 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uiParam2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read uiParam2 from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strParam.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateNotificationData(fmt.Errorf("Failed to read strParam from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateNotificationData(nil, packet, callID, uiType, uiParam1, uiParam2, strParam)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
