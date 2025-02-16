// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleInsertItem(packet nex.PacketInterface) {
	if protocol.InsertItem == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "PersistentStore::InsertItem not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	uiGroup := types.NewPrimitiveU32(0)
	strTag := types.NewString("")
	bufData := types.NewBuffer(nil)
	bReplace := types.NewPrimitiveBool(false)

	var err error

	err = uiGroup.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertItem(fmt.Errorf("Failed to read uiGroup from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strTag.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertItem(fmt.Errorf("Failed to read strTag from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = bufData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertItem(fmt.Errorf("Failed to read bufData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = bReplace.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.InsertItem(fmt.Errorf("Failed to read bReplace from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.InsertItem(nil, packet, callID, uiGroup, strTag, bufData, bReplace)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
