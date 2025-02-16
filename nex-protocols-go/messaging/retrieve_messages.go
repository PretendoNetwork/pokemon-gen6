// Package protocol implements the Messaging protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	messaging_types "github.com/PretendoNetwork/nex-protocols-go/v2/messaging/types"
)

func (protocol *Protocol) handleRetrieveMessages(packet nex.PacketInterface) {
	if protocol.RetrieveMessages == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Messaging::RetrieveMessages not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	recipient := messaging_types.NewMessageRecipient()
	lstMsgIDs := types.NewList[*types.PrimitiveU32]()
	lstMsgIDs.Type = types.NewPrimitiveU32(0)
	bLeaveOnServer := types.NewPrimitiveBool(false)

	var err error

	err = recipient.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RetrieveMessages(fmt.Errorf("Failed to read recipient from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = lstMsgIDs.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RetrieveMessages(fmt.Errorf("Failed to read lstMsgIDs from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = bLeaveOnServer.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.RetrieveMessages(fmt.Errorf("Failed to read bLeaveOnServer from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.RetrieveMessages(nil, packet, callID, recipient, lstMsgIDs, bLeaveOnServer)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
