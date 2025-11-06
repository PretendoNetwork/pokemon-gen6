package nex_subscription

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	subscription_types "github.com/PretendoNetwork/nex-protocols-go/v2/subscription/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
)

func CreateMySubscriptionData(err error, packet nex.PacketInterface, callID uint32, unk types.UInt32, param subscription_types.SubscriptionData, unk2 types.Bool) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	globals.Logger.Infof("CreateMySubscriptionData | unk: %d | unk2: %s", unk, unk2.String())

	param.PrincipalID = client.PID()
	globals.Timeline.CreateData(client.PID(), param, globals.DataTargets)
	globals.HandleSubscriptionChangeNotification(client.PID())

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodCreateMySubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
