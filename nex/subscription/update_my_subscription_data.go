package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription_types "github.com/PretendoNetwork/pokemon-gen6/nex/subscription/types"
)

func UpdateMySubscriptionData(err error, packet nex.PacketInterface, callID uint32, param subscription_types.SubscriptionData) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	globals.Timeline[uint32(client.PID())] = param.Data.Copy().(types.QBuffer)

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodUpdateMySubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
