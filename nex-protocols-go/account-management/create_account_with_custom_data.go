// Package protocol implements the Account Management protocol
package protocol

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
)

func (protocol *Protocol) handleCreateAccountWithCustomData(packet nex.PacketInterface) {
	if protocol.CreateAccountWithCustomData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "AccountManagement::CreateAccountWithCustomData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	strPrincipalName := types.NewString("")
	strKey := types.NewString("")
	uiGroups := types.NewPrimitiveU32(0)
	strEmail := types.NewString("")
	oPublicData := types.NewAnyDataHolder()
	oPrivateData := types.NewAnyDataHolder()

	var err error

	err = strPrincipalName.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read strPrincipalName from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strKey.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read strKey from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = uiGroups.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read uiGroups from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = strEmail.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read strEmail from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = oPublicData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read oPublicData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	err = oPrivateData.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.CreateAccountWithCustomData(fmt.Errorf("Failed to read oPrivateData from parameters. %s", err.Error()), packet, callID, nil, nil, nil, nil, nil, nil)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.CreateAccountWithCustomData(nil, packet, callID, strPrincipalName, strKey, uiGroups, strEmail, oPublicData, oPrivateData)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
