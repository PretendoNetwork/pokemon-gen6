package nex_subscription

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func GetSubscriptionData(err error, packet nex.PacketInterface, callID uint32, pids types.List[types.PID]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	subscriptionDataList := types.NewList[subscription_types.SubscriptionData]()

	for _, pid := range pids {
		globals.Logger.Infof("GetSubscriptionData pid: %d", pid)

		subscriptionData := subscription_types.NewSubscriptionData()

		data, err := globals.SubscriptionTimeline.GetData(pid)
		if err != nil {
			return nil, err
		}

		subscriptionData.PrincipalID = data.Data.PrincipalID
		subscriptionData.Unknown = data.Data.Unknown

		subscriptionDataList = append(subscriptionDataList, subscriptionData)
	}

	subscriptionDataList.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
