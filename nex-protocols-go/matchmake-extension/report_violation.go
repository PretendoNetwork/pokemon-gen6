// Package protocol implements the Matchmake Extension protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleReportViolation(packet nex.PacketInterface) {
	if protocol.ReportViolation == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "MatchmakeExtension::ReportViolation not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	pid := types.NewPID(0)
	userName := types.NewString("")
	violationCode := types.NewPrimitiveU32(0)

	var err error

	err = pid.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportViolation(fmt.Errorf("Failed to read pid from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = userName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportViolation(fmt.Errorf("Failed to read userName from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = violationCode.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.ReportViolation(fmt.Errorf("Failed to read violationCode from parameters. %s", err.Error()), packet, callID, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.ReportViolation(nil, packet, callID, pid, userName, violationCode)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
