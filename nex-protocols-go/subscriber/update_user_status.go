// Package protocol implements the Subscriber protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	subscriber_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscriber/types"
)

func (protocol *Protocol) handleUpdateUserStatus(packet nex.PacketInterface) {
	if protocol.UpdateUserStatus == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "Subscriber::UpdateUserStatus not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	unknown1 := types.NewList[*subscriber_types.Unknown]()
	unknown1.Type = subscriber_types.NewUnknown()
	unknown2 := types.NewList[*types.PrimitiveU8]()
	unknown2.Type = types.NewPrimitiveU8(0)

	var err error

	err = unknown1.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateUserStatus(fmt.Errorf("Failed to read unknown1 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = unknown2.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateUserStatus(fmt.Errorf("Failed to read unknown2 from parameters. %s", err.Error()), packet, callID, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateUserStatus(nil, packet, callID, unknown1, unknown2)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
