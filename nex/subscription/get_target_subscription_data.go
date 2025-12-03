package nex_subscription

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func GetTargetSubscriptionData(err error, packet nex.PacketInterface, callID uint32) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	rmcResponseStream := nex.NewByteStreamOut(endpoint.LibraryVersions(), endpoint.ByteStreamSettings())

	targetSubscriptionList := types.NewList[subscription_types.SubscriptionData]()

	if len(globals.SubscriptionTargets.GetTargets(client.PID())) > 0 {
		for _, target := range globals.SubscriptionTargets.GetTargets(client.PID()) {
			if !globals.SubscriptionTimeline.HasData(target) {
				continue
			}

			data, err := globals.SubscriptionTimeline.GetData(target)
			if err != nil {
				return nil, err
			}

			targetSubscriptionList = append(targetSubscriptionList, data.Data.Copy().(subscription_types.SubscriptionData))
		}
	}

	targetSubscriptionList.WriteTo(rmcResponseStream)

	rmcResponseBody := rmcResponseStream.Bytes()

	rmcResponse := nex.NewRMCSuccess(endpoint, rmcResponseBody)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodGetTargetSubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
