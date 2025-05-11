package nex_subscription

import (
	"encoding/hex"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription_types "github.com/PretendoNetwork/pokemon-gen6/nex/subscription/types"
)

func GetActivePlayerSubscriptionData(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	activePlayerSubscriptionList := types.NewList[subscription_types.ActivePlayerSubscriptionData]()

	for clientPID := range globals.Timeline {
		activePlayerSubscription := subscription_types.NewActivePlayerSubscriptionData()

		activePlayerSubscription.ByteUnk = 1
		activePlayerSubscription.SubscriptionData.OwnerPID = types.UInt32(clientPID)
		activePlayerSubscription.SubscriptionData.Data = globals.Timeline[clientPID]

		activePlayerSubscriptionList = append(activePlayerSubscriptionList, activePlayerSubscription)
	}

	activePlayerSubscriptionList.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	globals.Logger.Info(hex.EncodeToString(rmcResponseBody))

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetActivePlayerSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
