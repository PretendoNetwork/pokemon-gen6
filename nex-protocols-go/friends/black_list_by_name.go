// Package protocol implements the Friends QRV protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleBlackListByName(packet nex.PacketInterface) {
	if protocol.BlackListByName == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Friends::BlackListByName not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	strPlayerName := types.NewString("")
	uiDetails := types.NewPrimitiveU32(0)

	var err error

	err = strPlayerName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BlackListByName(fmt.Errorf("Failed to read strPlayerName from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uiDetails.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.BlackListByName(fmt.Errorf("Failed to read uiDetails from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.BlackListByName(nil, packet, callID, strPlayerName, uiDetails)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
