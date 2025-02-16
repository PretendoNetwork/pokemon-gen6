// Package protocol implements the DataStoreSuperMarioMaker protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleSetApplicationConfig(packet nex.PacketInterface) {
	if protocol.SetApplicationConfig == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "DataStoreSuperMarioMaker::SetApplicationConfig not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	applicationID := types.NewPrimitiveU32(0)
	key := types.NewPrimitiveU32(0)
	value := types.NewPrimitiveS32(0)

	var err error

	err = applicationID.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetApplicationConfig(fmt.Errorf("Failed to read applicationID from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = key.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetApplicationConfig(fmt.Errorf("Failed to read key from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = value.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.SetApplicationConfig(fmt.Errorf("Failed to read value from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.SetApplicationConfig(nil, packet, callID, applicationID, key, value)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
