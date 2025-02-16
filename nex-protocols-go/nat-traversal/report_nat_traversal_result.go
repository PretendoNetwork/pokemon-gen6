// Package protocol implements the NAT Traversal protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleReportNATTraversalResult(packet nex.PacketInterface) {
	if protocol.ReportNATTraversalResult == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "NATTraversal::ReportNATTraversalResult not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	endpoint := packet.Sender().Endpoint()
	natTraversalVersion := endpoint.LibraryVersions().NATTraversal

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	cid := types.NewPrimitiveU32(0)
	result := types.NewPrimitiveBool(false)
	rtt := types.NewPrimitiveU32(0)

	var err error

	err = cid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATTraversalResult(fmt.Errorf("Failed to read cid from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = result.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportNATTraversalResult(fmt.Errorf("Failed to read result from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	// TODO - Is this the right version?
	if natTraversalVersion.GreaterOrEqual("3.0.0") {
		err = rtt.ExtractFrom(parametersStream)
		if err != nil {
			_, rmcError := protocol.ReportNATTraversalResult(fmt.Errorf("Failed to read rtt from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
			if rmcError != nil {
				globals.RespondError(packet, ProtocolID, rmcError)
			}

			return
		}
	}

	rmcMessage, rmcError := protocol.ReportNATTraversalResult(nil, packet, callID, cid, result, rtt)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
