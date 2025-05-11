package nex_subscription

import (
	"encoding/hex"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription_types "github.com/PretendoNetwork/pokemon-gen6/nex/subscription/types"
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
		if globals.Timeline[pid] != nil {
			friendSubscriptionData := subscription_types.NewSubscriptionData()

			friendSubscriptionData.OwnerPID = types.UInt32(pid)
			friendSubscriptionData.Data = globals.Timeline[pid]

			friendSubscriptionDataList = append(friendSubscriptionDataList, friendSubscriptionData)
		}
	}

	friendSubscriptionDataList.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	globals.Logger.Info(hex.EncodeToString(rmcResponseBody))

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetFriendSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
