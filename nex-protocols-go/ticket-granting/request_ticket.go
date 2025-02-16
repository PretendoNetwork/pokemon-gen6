// Package protocol implements the Ticket Granting protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleRequestTicket(packet nex.PacketInterface) {
	if protocol.RequestTicket == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "TicketGranting::RequestTicket not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	idSource := types.NewPID(0)
	idTarget := types.NewPID(0)

	var err error

	err = idSource.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestTicket(fmt.Errorf("Failed to read idSource from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = idTarget.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RequestTicket(fmt.Errorf("Failed to read idTarget from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RequestTicket(nil, packet, callID, idSource, idTarget)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
