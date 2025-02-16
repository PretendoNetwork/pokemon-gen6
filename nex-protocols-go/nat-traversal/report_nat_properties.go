// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleReportNATProperties(packet nex.PacketInterface) {
	if protocol.ReportNATProperties == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::ReportNATProperties not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	natmapping := types.NewPrimitiveU32(0)
	natfiltering := types.NewPrimitiveU32(0)
	rtt := types.NewPrimitiveU32(0)

	var err error

	err = natmapping.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATProperties(fmt.Errorf("Failed to read natmapping from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = natfiltering.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATProperties(fmt.Errorf("Failed to read natfiltering from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = rtt.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATProperties(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportNATProperties(nil, packet, callID, natmapping, natfiltering, rtt)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
