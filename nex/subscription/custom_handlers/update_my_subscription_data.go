package nex_subscription_custom_handlers

import (
	"fmt"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	subscription_types "github.com/PretendoNetwork/pokemon-gen6/nex/subscription/types"
)

func (protocol *Protocol) handleUpdateMySubscriptionData(packet nex.PacketInterface) {
	if protocol.UpdateMySubscriptionData == nil {
		err := nex.NewError(nex.ResultCodes.Core.NotImplemented, "SubscriptionProtocol::UpdateMySubscriptionData not implemented")

		globals.Logger.Warning(err.Message)
		globals.RespondError(packet, ProtocolID, err)

		return
	}

	request := packet.RMCMessage()
	callID := request.CallID
	parameters := request.Parameters
	endpoint := packet.Sender().Endpoint()
	parametersStream := nex.NewByteStreamIn(parameters, endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	var param subscription_types.SubscriptionData

	err := param.ExtractFrom(parametersStream)
	if err != nil {
		_, rmcError := protocol.UpdateMySubscriptionData(fmt.Errorf("Failed to read param from parameters. %s", err.Error()), packet, callID, param)
		if rmcError != nil {
			globals.RespondError(packet, ProtocolID, rmcError)
		}

		return
	}

	rmcMessage, rmcError := protocol.UpdateMySubscriptionData(nil, packet, callID, param)
	if rmcError != nil {
		globals.RespondError(packet, ProtocolID, rmcError)
		return
	}

	globals.Respond(packet, rmcMessage)
}
