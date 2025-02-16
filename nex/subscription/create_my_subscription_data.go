package nex_subscription

import (
	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	"github.com/PretendoNetwork/pokemon-gen6/globals"
	subscription "github.com/PretendoNetwork/nex-protocols-go/v2/subscription"
)

func CreateMySubscriptionData(err error, packet nex.PacketInterface, callID uint32, unk *types.PrimitiveU64, content []byte) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.Core.InvalidArgument, err.Error())
	}

	client := packet.Sender()

	endpoint := client.Endpoint().(*nex.PRUDPEndPoint)

	globals.Timeline[client.PID().LegacyValue()] = content

	rmcResponse := nex.NewRMCSuccess(endpoint, nil)
	rmcResponse.ProtocolID = subscription.ProtocolID
	rmcResponse.MethodID = subscription.MethodCreateMySubscriptionData
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
