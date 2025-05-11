package nex_subscription

import (
	"encoding/hex"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription_types "github.com/PretendoNetwork/pokemon-gen6/nex/subscription/types"
)

func GetSubscriptionData(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.UInt32]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	subscriptionDataList := types.NewList[subscription_types.SubscriptionData]()

	for _, pid := range pids {
		globals.Logger.Infof("%d", pid)

		subscriptionData := subscription_types.NewSubscriptionData()

		subscriptionData.OwnerPID = pid
		subscriptionData.Data = globals.Timeline[uint32(pid)]

		subscriptionDataList = append(subscriptionDataList, subscriptionData)
	}

	subscriptionDataList.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	globals.Logger.Info(hex.EncodeToString(rmcResponseBody))

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
