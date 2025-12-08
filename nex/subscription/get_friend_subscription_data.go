package nex_subscription

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func GetFriendSubscriptionData(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	friendPids := globals.GetUserFriendPIDs(uint32(packet.Sender().PID()))

	friendSubscriptionDataList := types.NewList[subscription_types.SubscriptionData]()

	for _, pid := range friendPids {
		if globals.SubscriptionTimeline.HasData(types.PID(pid)) {
			friendSubscriptionData := subscription_types.NewSubscriptionData()

			data, err := globals.SubscriptionTimeline.GetData(types.PID(pid))
			if err != nil {
				return nil, err
			}

			friendSubscriptionData.PrincipalID = data.Data.PrincipalID
			friendSubscriptionData.Unknown = data.Data.Unknown.Copy().(types.QBuffer)

			friendSubscriptionDataList = append(friendSubscriptionDataList, friendSubscriptionData)
		}
	}

	friendSubscriptionDataList.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetFriendSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
