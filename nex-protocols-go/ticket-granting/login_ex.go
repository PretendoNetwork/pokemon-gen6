// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleLoginEx(packet nex.PacketInterface) {
	if protocol.LoginEx == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "TicketGranting::LoginEx not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	strUserName := types.NewString("")
	oExtraData := types.NewAnyDataHolder()

	var err error

	err = strUserName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginEx(fmt.Errorf("Failed to read strUserName from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = oExtraData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.LoginEx(fmt.Errorf("Failed to read oExtraData from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.LoginEx(nil, packet, callID, strUserName, oExtraData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
